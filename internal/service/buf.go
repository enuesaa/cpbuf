package service

import (
	"path/filepath"
	"slices"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/cpbuf/internal/repository"
)

func NewBufSrv(repos repository.Repos) BufSrv {
	return BufSrv{
		repos: repos,
	}
}

type BufSrv struct {
	repos repository.Repos
}

func (srv *BufSrv) GetBufDirPath() (string, error) {
	homedir, err := srv.repos.Fs.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".cpbuf"), nil
}

func (srv *BufSrv) IsBufDirExist() bool {
	path, err := srv.GetBufDirPath()
	if err != nil {
		return false
	}
	return srv.repos.Fs.IsExist(path)
}

func (srv *BufSrv) CreateBufDir() error {
	if srv.IsBufDirExist() {
		return nil
	}
	path, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.CreateDir(path)
}

func (srv *BufSrv) DeleteBufDir() error {
	if !srv.IsBufDirExist() {
		return nil
	}
	path, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	return srv.repos.Fs.Remove(path)
}

func (srv *BufSrv) Buffer(filename string) error {
	isDir, err := srv.repos.Fs.IsDir(filename)
	if err != nil {
		return err
	}
	if isDir {
		return srv.BufferDir(filename)
	}
	return srv.BufferFile(filename)
}

func (srv *BufSrv) BufferFile(filename string) error {
	bufferPath, err := srv.GetBufferPath(filename)
	if err != nil {
		return err
	}
	workPath, err := srv.GetWorkPath(filename)
	if err != nil {
		return err
	}
	return srv.repos.Fs.CopyFile(workPath, bufferPath)
}

func (srv *BufSrv) BufferDir(dirname string) error {
	files, err := srv.ListFilesRecursively(dirname)
	if err != nil {
		return err
	}
	for _, filename := range files {
		bufferPath, err := srv.GetBufferPath(filename)
		if err != nil {
			return err
		}
		if err := srv.repos.Fs.CreateDir(filepath.Dir(bufferPath)); err != nil {
			return err
		}
		if err := srv.Buffer(filename); err != nil {
			return err
		}
	}
	return nil
}

func (srv *BufSrv) GetBufferPath(filename string) (string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(bufDirPath, filename), nil
}

func (srv *BufSrv) GetWorkPath(filename string) (string, error) {
	workDirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(workDirPath, filename), nil
}

func (srv *BufSrv) PasteFile(path string) error {
	return nil
}

//Deprecated: use PasteFile instead.
func (srv *BufSrv) PasteFileToWorkDir(filename string) error {
	registryPath, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	filePathInBufDir := filepath.Join(registryPath, filename)

	isDir, err := srv.repos.Fs.IsDir(filePathInBufDir)
	if err != nil {
		return err
	}
	if isDir {
		files, err := srv.ListFilesRecursively(filePathInBufDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			relpath, err := filepath.Rel(registryPath, file)
			if err := srv.repos.Fs.CreateDir(filepath.Dir(relpath)); err != nil {
				return err
			}
			isfiledir, err := srv.repos.Fs.IsDir(file)
			if err != nil {
				return err
			}
			if isfiledir {
				continue
			}
			if err := srv.repos.Fs.CopyFile(file, relpath); err != nil {
				return err
			}
		}
		return nil
	}
	return srv.repos.Fs.CopyFile(filePathInBufDir, filename)
}

func (srv *BufSrv) ListFilesRecursively(path string) ([]string, error) {
	files, err := srv.repos.Fs.ListFiles(path)
	if err != nil {
		return make([]string, 0), err
	}

	list := make([]string, 0)
	for _, filename := range files {
		fpath := filepath.Join(path, filename)
		list = append(list, fpath)
		isDir, err := srv.repos.Fs.IsDir(fpath)
		if err != nil {
			return make([]string, 0), err
		}
		if isDir {
			innerList, err := srv.ListFilesRecursively(fpath)
			if err != nil {
				return make([]string, 0), err
			}
			list = append(list, innerList...)
		}
	}

	return list, nil
}

func (srv *BufSrv) RemoveFileInWorkDir(filename string) error {
	workdirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return err
	}

	return srv.repos.Fs.Remove(filepath.Join(workdirPath, filename))
}

func (srv *BufSrv) ListFilesInWorkDir() ([]string, error) {
	workdirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return make([]string, 0), err
	}

	return srv.repos.Fs.ListFiles(workdirPath)
}

func (srv *BufSrv) ListFilesInBufDir() ([]string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return make([]string, 0), err
	}
	return srv.repos.Fs.ListFiles(bufDirPath)
}

func (srv *BufSrv) SelectFileWithPrompt() string {
	filename := srv.repos.Prompt.StartSelectPrompt("filename: ", func(in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)

		files, _ := srv.repos.Fs.ListFiles(".")
		for _, filename := range files {
			suggests = append(suggests, prompt.Suggest{Text: filename})
		}

		return prompt.FilterHasPrefix(suggests, in.Text, false)
	})

	return filename
}

func (srv *BufSrv) ListConflictedFilenames() ([]string, error) {
	workdirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return make([]string, 0), err
	}

	workdirFilenames, err := srv.repos.Fs.ListFiles(workdirPath)
	if err != nil {
		return make([]string, 0), err
	}

	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return make([]string, 0), err
	}
	bufDirFilenames, err := srv.repos.Fs.ListFiles(bufDirPath)

	duplicates := make([]string, 0)
	for _, filename := range bufDirFilenames {
		if slices.Contains(workdirFilenames, filename) {
			duplicates = append(duplicates, filename)
		}
	}

	return duplicates, nil
}
