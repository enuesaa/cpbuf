package task

import (
	"github.com/enuesaa/cpbuf/internal/repository"
)

func NewWorkfile(repos repository.Repos, filename string) Workfile {
	return Workfile{
		repos: repos,
		filename: filename,
	}
}

type Workfile struct {
	repos repository.Repos
	filename string
}

func (f *Workfile) GetWorkPath() string {
	return ""
}

func (f *Workfile) CopyToBufDir() error {
	return nil
}
