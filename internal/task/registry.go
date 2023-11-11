package task

import (
	"github.com/enuesaa/cpbuf/internal/repository"
)

func NewRegistry(repos repository.Repos) Registry {
	return Registry {
		repos: repos,
	}
}

type Registry struct {
	repos repository.Repos
}

func (srv *Registry) CreateBufDir() error {
	return nil
}
func (srv *Registry) DeleteBufDir() error {
	return nil
}

func (srv *Registry) ListFilesInBufDir() ([]Bufferfile, error) {
	return make([]Bufferfile, 0), nil
}

func (srv *Registry) ListFilesInWorkDir() ([]Workfile, error) {
	return make([]Workfile, 0), nil
}

func (srv *Registry) ListConflictedFilenames() ([]string, error) {
	return make([]string, 0), nil
}
