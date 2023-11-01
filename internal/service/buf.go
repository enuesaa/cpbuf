package service

import (
	"path/filepath"
	"slices"

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

func (srv *BufSrv) PasteFilesToWorkDir() error {
	bufDirPath, err := srv.repos.Fshome.GetResgistryPath(srv.getBufDirName())
	if err != nil {
		return err
	}
	filenames, err := srv.repos.Fshome.ListFiles(bufDirPath)
	for _, filename := range filenames {
		if err := srv.CopyFileToBufDir(filename); err != nil {
			return err
		}
	}
	return nil
}

func (srv *BufSrv) ExtractSameFilenamesInWorkDir() ([]string, error) {
	workdirPath, err := srv.repos.Fshome.Workdir()
	if err != nil {
		return make([]string, 0), err
	}

	workdirFilenames, err := srv.repos.Fshome.ListFiles(workdirPath)
	if err != nil {
		return make([]string, 0), err
	}

	bufDirPath, err := srv.repos.Fshome.GetResgistryPath(srv.getBufDirName())
	bufDirFilenames, err := srv.repos.Fshome.ListFiles(bufDirPath)

	duplicates := make([]string, 0)
	for _, filename := range bufDirFilenames {
		if slices.Contains(workdirFilenames, filename) {
			duplicates = append(duplicates, filename)
		}
	}

	return duplicates, nil
}

func (srv *BufSrv) ListFilesInWorkDir() ([]string, error) {
	workdirPath, err := srv.repos.Fshome.Workdir()
	if err != nil {
		return make([]string, 0), err
	}

	return srv.repos.Fshome.ListFiles(workdirPath)
}
