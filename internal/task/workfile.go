package task

import (
	"path/filepath"
	"strings"

	"github.com/enuesaa/cpbuf/internal/repository"
)

func NewWorkfile(repos repository.Repos, path string, workDir string, bufferDir string) Workfile {
	return Workfile {
		repos: repos,
		path: path,
		workDir: workDir,
		bufferDir: bufferDir,
	}
}

type Workfile struct {
	repos repository.Repos
	path string
	workDir string
	bufferDir string
}

func (f *Workfile) IsDir() (bool, error) {
	path, err := f.GetWorkPath()
	if err != nil {
		return false, err
	}
	return f.repos.Fs.IsDir(path)
}

func (f *Workfile) GetFilename() string {
	return strings.TrimPrefix(f.path, f.workDir + "/")
}

func (f *Workfile) GetBufferPath() string {
	return filepath.Join(f.bufferDir, f.GetFilename())
}

func (f *Workfile) GetWorkPath() (string, error) {
	return filepath.Join(f.workDir, f.GetFilename()), nil
}
