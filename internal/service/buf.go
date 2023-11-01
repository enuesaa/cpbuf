package service

import "github.com/enuesaa/cpbuf/internal/repository"

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


