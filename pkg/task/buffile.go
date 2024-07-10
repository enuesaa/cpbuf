package task

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/cpbuf/pkg/repository"
)

type Buffile struct {
	repos     repository.Repos
	path      string
	bufferDir string
	workDir   string
}

func (f *Buffile) IsDir() (bool, error) {
	path := f.GetBufferPath()
	return f.repos.Fs.IsDir(path)
}

func (f *Buffile) GetFilename() string {
	return strings.TrimPrefix(f.path, f.bufferDir+"/")
}

func (f *Buffile) GetBufferPath() string {
	return filepath.Join(f.bufferDir, f.GetFilename())
}

func (f *Buffile) IsTopLevel() bool {
	return !strings.Contains(f.GetFilename(), "/")
}

func (f *Buffile) IsSameFileExistInWorkDir() (bool, error) {
	workDirPath, err := f.repos.Fs.WorkDir()
	if err != nil {
		return false, err
	}
	workPath := filepath.Join(workDirPath, f.GetFilename())
	return f.repos.Fs.IsExist(workPath), nil
}

func (f *Buffile) GetWorkPath() (string, error) {
	return filepath.Join(f.workDir, f.GetFilename()), nil
}

func (f *Buffile) GetBufferedDate() string {
	modtime, err := f.repos.Fs.GetModTime(f.GetBufferPath())
	if err != nil {
		return ""
	}
	return modtime.Format("2006/01/02")
}
