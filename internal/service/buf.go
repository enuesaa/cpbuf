package service

import (
	"path/filepath"

	"github.com/enuesaa/cpbuf/internal/repository"
)

type BufSrv struct {
	repos repository.Repos
}
func NewBufSrv(repos repository.Repos) BufSrv {
	return BufSrv{
		repos: repos,
	}
}

func (srv *BufSrv) getBufDirName() string {
	return ".cpbuf"
}

func (srv *BufSrv) CreateBufDir() error {
	if srv.repos.Fshome.IsRegistryExist(srv.getBufDirName()) {
		return nil
	}

	return srv.repos.Fshome.CreateRegistry(srv.getBufDirName())
}

func (srv *BufSrv) DeleteBufDir() error {
	if srv.repos.Fshome.IsRegistryExist(srv.getBufDirName()) {
		return srv.repos.Fshome.DeleteRegistry(srv.getBufDirName())
	}
	return nil
}

func (srv *BufSrv) CopyFileToBufDir(filename string) error {
	registryPath, err := srv.repos.Fshome.GetResgistryPath(srv.getBufDirName())
	if err != nil {
		return err
	}
	dstPath := filepath.Join(registryPath, filename)
	return srv.repos.Fshome.CopyFile(filename, dstPath)
}

