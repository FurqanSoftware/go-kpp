package kpp

import (
	"io"
	"io/fs"
	"path"
	"strings"
)

type Statement struct {
	fs fs.FS

	name string

	language string
	typ      string
}

func newStatement(fs fs.FS, name string) (stmt Statement) {
	stmt.fs = fs
	stmt.name = name
	stmt.language = "en"
	parts := strings.Split(name, ".")
	if len(parts) == 3 {
		stmt.language = parts[1]
	}
	stmt.typ = parts[len(parts)-1]
	return
}

func (s Statement) Language() string {
	return s.language
}

// Type returns tex, md, or pdf.
func (s Statement) Type() string {
	return s.typ
}

func (s Statement) Content() (io.ReadCloser, error) {
	return s.fs.Open(path.Join("problem_statement", s.name))
}
