package kpp

import (
	"io"
	"io/fs"
	"path"
)

type Submission struct {
	fs fs.FS

	typ  string
	name string
}

// Type returns tex, md, or pdf.
func (s Submission) Type() string {
	return s.typ
}

func (s Submission) Source() (io.ReadCloser, error) {
	return s.fs.Open(path.Join("submissions", s.typ, s.name))
}
