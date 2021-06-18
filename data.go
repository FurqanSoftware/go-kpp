package kpp

import (
	"io"
	"io/fs"
	"path"
)

type TestData struct {
	fs fs.FS

	group *TestGroup
	base  string

	illusExt string
}

func (d TestData) Group() string {
	return d.group.Path()
}

func (d TestData) Input() (io.ReadCloser, error) {
	return d.fs.Open(path.Join("data", d.group.Path(), d.base+".in"))
}

func (d TestData) Answer() (io.ReadCloser, error) {
	return d.fs.Open(path.Join("data", d.group.Path(), d.base+".ans"))
}

func (d TestData) readText(name string) (string, error) {
	f, err := d.fs.Open(path.Join("data", d.group.Path(), name))
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	return string(b), nil
}

func (d TestData) Hint() (string, error) {
	return d.readText(d.base + ".hint")
}

func (d TestData) Description() (string, error) {
	return d.readText(d.base + ".desc")
}

func (d TestData) Illustration() (io.ReadCloser, string, error) {
	if d.illusExt == "" {
		return nil, "", nil
	}
	f, err := d.fs.Open(path.Join("data", d.group.Path(), d.base+"."+d.illusExt))
	return f, d.illusExt, err
}

type TestGroup struct {
	parent *TestGroup

	name string
}

func (g TestGroup) Path() string {
	if g.parent == nil {
		return g.name
	}
	return path.Join(g.parent.Path(), g.name)
}
