package task

import (
	"github.com/enuesaa/cpbuf/internal/repository"
)

func NewBufferfile(repos repository.Repos, filename string) Bufferfile {
	return Bufferfile{
		repos: repos,
		filename: filename,
	}
}

type Bufferfile struct {
	repos repository.Repos
	filename string
}

func (f *Bufferfile) GetBufferPath() string {
	return ""
}

func (f *Bufferfile) CopyToWorkDir() error {
	return nil
}
