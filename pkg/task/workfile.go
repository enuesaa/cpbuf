package task

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/enuesaa/cpbuf/pkg/repository"
)

func NewWorkfile(repos repository.Repos, path string, bufferDir string, workDir string) Workfile {
	return Workfile{
		repos:     repos,
		path:      path,
		workDir:   workDir,
		bufferDir: bufferDir,
	}
}

type Workfile struct {
	repos     repository.Repos
	path      string
	workDir   string
	bufferDir string
}

func (f *Workfile) IsDir() (bool, error) {
	path, err := f.GetWorkPath()
	if err != nil {
		return false, err
	}
	isDir, err := f.repos.Fs.IsDir(path)
	if err != nil {
		if isBrokenSymlink, e := f.IsBrokenSymlink(); e != nil || !isBrokenSymlink {
			return false, err
		}
		isDir = false
	}
	return isDir, err
}

func (f *Workfile) GetFilename() string {
	return strings.TrimPrefix(f.path, f.workDir+"/")
}

func (f *Workfile) GetBufferPath() string {
	return filepath.Join(f.bufferDir, f.GetFilename())
}

func (f *Workfile) GetWorkPath() (string, error) {
	return filepath.Join(f.workDir, f.GetFilename()), nil
}

func (f *Workfile) CheckExist() error {
	path, err := f.GetWorkPath()
	if err != nil {
		return err
	}
	if !f.repos.Fs.IsExist(path) {
		return fmt.Errorf("file %s not found", f.GetFilename())
	}
	return nil
}

func (f *Workfile) IsBrokenSymlink() (bool, error) {
	path, err := f.GetWorkPath()
	if err != nil {
		return false, err
	}
	return f.repos.Fs.IsBrokenSymlink(path)
}
