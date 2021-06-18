package kpp

import (
	"io/fs"
	"path"
	"strings"

	"github.com/goccy/go-yaml"
)

type Package struct {
	fs fs.FS

	metadata Metadata
}

func Open(fs fs.FS) (*Package, error) {
	pkg := Package{fs: fs}
	err := pkg.readMetadata()
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (p Package) Metadata() Metadata {
	return p.metadata
}

func (p Package) Statements() ([]Statement, error) {
	d, err := fs.ReadDir(p.fs, "problem_statement")
	if err != nil {
		return nil, err
	}
	stmts := []Statement{}
	for _, e := range d {
		if !strings.HasPrefix(e.Name(), "problem.") {
			continue
		}
		ext := path.Ext(e.Name())
		if ext != ".tex" && ext != ".md" && ext != ".pdf" {
			continue
		}

		stmts = append(stmts, newStatement(p.fs, e.Name()))
	}
	return stmts, nil
}

func (p *Package) TestData() ([]TestData, error) {
	tests := []TestData{}
	var g *TestGroup
	err := fs.WalkDir(p.fs, "data", func(name string, d fs.DirEntry, err error) error {
		if name == "data" {
			return nil
		}
		name = strings.TrimPrefix(name, "data/")

		if d.IsDir() {
			for g != nil && !strings.HasPrefix(name, g.Path()) {
				g = g.parent
			}
			g = &TestGroup{
				parent: g,
				name:   path.Base(name),
			}
			return nil
		}

		if path.Ext(name) != ".in" {
			return nil
		}

		test := TestData{
			fs:    p.fs,
			group: g,
			name:  strings.TrimSuffix(path.Base(name), ".in"),
		}
		fi, err := fs.Stat(p.fs, path.Join("data", g.Path(), test.name+".in"))
		if err != nil {
			return err
		}
		test.inputSize = fi.Size()
		fi, err = fs.Stat(p.fs, path.Join("data", g.Path(), test.name+".ans"))
		if err != nil {
			return err
		}
		test.answerSize = fi.Size()
		for _, ext := range []string{".png", ".jpg", ".jpeg", ".svg"} {
			_, err := fs.Stat(p.fs, path.Join("data", g.Path(), test.name+ext))
			if err == nil {
				test.illusExt = ext
				break
			}
		}

		tests = append(tests, test)
		return nil
	})
	return tests, err
}

func (p *Package) readMetadata() error {
	f, err := p.fs.Open("problem.yaml")
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(&p.metadata)
}
