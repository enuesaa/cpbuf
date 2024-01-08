package task

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/cpbuf/pkg/repository"
)

func NewBufferfile(repos repository.Repos, path string, bufferDir string, workDir string) Bufferfile {
	return Bufferfile{
		repos:     repos,
		path:      path,
		bufferDir: bufferDir,
		workDir:   workDir,
	}
}

type Bufferfile struct {
	repos     repository.Repos
	path      string
	bufferDir string
	workDir   string
}

func (f *Bufferfile) IsDir() (bool, error) {
	path := f.GetBufferPath()
	return f.repos.Fs.IsDir(path)
}

func (f *Bufferfile) GetFilename() string {
	return strings.TrimPrefix(f.path, f.bufferDir+"/")
}

func (f *Bufferfile) GetBufferPath() string {
	return filepath.Join(f.bufferDir, f.GetFilename())
}

func (f *Bufferfile) IsTopLevel() bool {
	return !strings.Contains(f.GetFilename(), "/")
}

func (f *Bufferfile) IsSameFileExistInWorkDir() (bool, error) {
	workDirPath, err := f.repos.Fs.WorkDir()
	if err != nil {
		return false, err
	}
	workPath := filepath.Join(workDirPath, f.GetFilename())
	return f.repos.Fs.IsExist(workPath), nil
}

func (f *Bufferfile) GetWorkPath() (string, error) {
	return filepath.Join(f.workDir, f.GetFilename()), nil
}
