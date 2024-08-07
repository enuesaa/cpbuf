package task

import (
	"path/filepath"
	"slices"

	"github.com/enuesaa/cpbuf/pkg/repository"
)

func NewRegistry(repos repository.Repos) Registry {
	return Registry{
		repos: repos,
	}
}

type Registry struct {
	repos repository.Repos
}

func (srv *Registry) GetBufDirPath() (string, error) {
	homedir, err := srv.repos.Fs.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".cpbuf"), nil
}

func (srv *Registry) IsBufDirExist() bool {
	path, err := srv.GetBufDirPath()
	if err != nil {
		return false
	}
	return srv.repos.Fs.IsExist(path)
}

func (srv *Registry) CreateBufDir() error {
	if srv.IsBufDirExist() {
		return nil
	}
	path, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.CreateDir(path)
}

func (srv *Registry) DeleteBufDir() error {
	if !srv.IsBufDirExist() {
		return nil
	}
	path, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.Remove(path)
}

func (srv *Registry) ListFilesInBufDir() ([]Buffile, error) {
	bufDir, err := srv.GetBufDirPath()
	if err != nil {
		return make([]Buffile, 0), err
	}
	files, err := srv.ListFilesRecursively(bufDir)
	if err != nil {
		return make([]Buffile, 0), err
	}
	workDir, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return make([]Buffile, 0), err
	}
	list := make([]Buffile, 0)
	for _, file := range files {
		list = append(list, Buffile{
			repos: srv.repos,
			path: file,
			bufferDir: bufDir,
			workDir: workDir,
		})
	}
	return list, nil
}

func (srv *Registry) GetWorkfile(filename string) Workfile {
	workDir, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return Workfile{}
	}
	bufDir, err := srv.GetBufDirPath()
	if err != nil {
		return Workfile{}
	}
	workfile := NewWorkfile(srv.repos, filename, bufDir, workDir)
	return workfile
}

func (srv *Registry) ListFilesInWorkDir() ([]Workfile, error) {
	workDir, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return make([]Workfile, 0), err
	}
	files, err := srv.ListFilesRecursively(workDir)
	if err != nil {
		return make([]Workfile, 0), err
	}
	bufDir, err := srv.GetBufDirPath()
	if err != nil {
		return make([]Workfile, 0), err
	}

	list := make([]Workfile, 0)
	for _, file := range files {
		list = append(list, NewWorkfile(srv.repos, file, bufDir, workDir))
	}
	return list, nil
}

func (srv *Registry) CopyToBufDir(workfile Workfile) error {
	bufferPath := workfile.GetBufferPath()
	isDir, err := workfile.IsDir()
	if err != nil {
		return err
	}
	if isDir {
		if err := srv.repos.Fs.CreateDir(bufferPath); err != nil {
			return err
		}
		return nil
	}
	workPath, err := workfile.GetWorkPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.CopyFile(workPath, bufferPath)
}

func (srv *Registry) CopyToWorkDir(buffile Buffile) error {
	workPath, err := buffile.GetWorkPath()
	if err != nil {
		return err
	}
	isDir, err := buffile.IsDir()
	if err != nil {
		return err
	}
	if isDir {
		if err := srv.repos.Fs.CreateDir(workPath); err != nil {
			return err
		}
		return nil
	}
	bufferPath := buffile.GetBufferPath()
	return srv.repos.Fs.CopyFile(bufferPath, workPath)
}

func (srv *Registry) ListFilesRecursively(path string) ([]string, error) {
	files, err := srv.repos.Fs.ListFiles(path)
	if err != nil {
		return make([]string, 0), err
	}

	for _, fpath := range files {
		isDir, err := srv.repos.Fs.IsDir(fpath)
		if err != nil {
			if isBrokenSymlink, e := srv.repos.Fs.IsBrokenSymlink(fpath); e != nil || !isBrokenSymlink {
				return make([]string, 0), err
			}
			isDir = false
		}
		if isDir {
			innerList, err := srv.ListFilesRecursively(fpath)
			if err != nil {
				return make([]string, 0), err
			}
			files = append(files, innerList...)
		}
	}

	slices.Sort(files)
	return slices.Compact(files), nil
}

func (srv *Registry) RemoveFileInWorkDir(filename string) error {
	workDir, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return err
	}
	workPath := filepath.Join(workDir, filename)
	return srv.repos.Fs.Remove(workPath)
}

func (srv *Registry) RemoveFileInBufDir(filename string) error {
	bufDir, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	bufFilePath := filepath.Join(bufDir, filename)
	return srv.repos.Fs.Remove(bufFilePath)
}
