package service

import (
	"path/filepath"
	"slices"
	"strings"

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

func (srv *BufSrv) GetBufferPath(filename string) (string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(filename, bufDirPath) {
		return filename, nil
	}
	return filepath.Join(bufDirPath, filename), nil
}

func (srv *BufSrv) GetWorkPath(filename string) (string, error) {
	workDirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(filename, workDirPath) {
		return filename, nil
	}
	return filepath.Join(workDirPath, filename), nil
}

func (srv *BufSrv) ConvertWorkPathToFilename(workPath string) (string, error) {
	workDirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(workPath, workDirPath + "/"), nil
}

func (srv *BufSrv) ConvertBufferPathToFilename(bufferPath string) (string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(bufferPath, bufDirPath + "/"), nil
}

func (srv *BufSrv) Buffer(filename string) error {
	workPath, err := srv.GetWorkPath(filename)
	if err != nil {
		return err
	}
	isDir, err := srv.repos.Fs.IsDir(workPath)
	if err != nil {
		return err
	}
	if isDir {
		return srv.BufferDir(workPath)
	}
	bufferPath, err := srv.GetBufferPath(filename)
	if err != nil {
		return err
	}
	return srv.repos.Fs.CopyFile(workPath, bufferPath)
}

func (srv *BufSrv) BufferDir(workPath string) error {
	files, err := srv.ListFilesRecursively(workPath)
	if err != nil {
		return err
	}
	for _, fpath := range files {
		filename, err := srv.ConvertWorkPathToFilename(fpath)
		if err != nil {
			return err
		}
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

func (srv *BufSrv) Paste(filename string) error {
	bufferPath, err := srv.GetBufferPath(filename)
	if err != nil {
		return err
	}
	isDir, err := srv.repos.Fs.IsDir(bufferPath)
	if err != nil {
		return err
	}
	if isDir {
		return srv.PasteDir(bufferPath)
	}
	workPath, err := srv.GetWorkPath(filename)
	if err != nil {
		return err
	}
	return srv.repos.Fs.CopyFile(bufferPath, workPath)
}

func (srv *BufSrv) PasteDir(bufferPath string) error {
	files, err := srv.ListFilesRecursively(bufferPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename, err := srv.ConvertBufferPathToFilename(file)
		if err != nil {
			return err
		}
		workPath, err := srv.GetWorkPath(filename)
		if err != nil {
			return err
		}
		if err := srv.repos.Fs.CreateDir(filepath.Dir(workPath)); err != nil {
			return err
		}
		if err := srv.Paste(filename); err != nil {
			return err
		}
	}
	return nil
}

func (srv *BufSrv) ListFilesRecursively(path string) ([]string, error) {
	files, err := srv.repos.Fs.ListFiles(path)
	if err != nil {
		return make([]string, 0), err
	}

	for _, fpath := range files {
		isDir, err := srv.repos.Fs.IsDir(fpath)
		if err != nil {
			return make([]string, 0), err
		}
		if isDir {
			innerList, err := srv.ListFilesRecursively(fpath)
			if err != nil {
				return make([]string, 0), err
			}
			files = append(files, innerList...)
		}
	}

	return files, nil
}

func (srv *BufSrv) RemoveFileInWorkDir(filename string) error {
	workPath, err := srv.GetWorkPath(filename)
	if err != nil {
		return err
	}
	return srv.repos.Fs.Remove(workPath)
}

func (srv *BufSrv) ListFilesInWorkDir() ([]string, error) {
	workDirPath, err := srv.repos.Fs.WorkDir()
	if err != nil {
		return make([]string, 0), err
	}
	return srv.repos.Fs.ListFiles(workDirPath)
}

func (srv *BufSrv) ListFilesInBufDir() ([]string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return make([]string, 0), err
	}
	files, err := srv.repos.Fs.ListFiles(bufDirPath)
	if err != nil {
		return make([]string, 0), err
	}
	list := make([]string, 0)
	for _, path := range files {
		filename, err := srv.ConvertBufferPathToFilename(path)
		if err != nil {
			return make([]string, 0), err
		}
		list = append(list, filename)
	}
	return list, nil
}

func (srv *BufSrv) SelectFileWithPrompt() string {
	filename := srv.repos.Prompt.StartSelectPrompt("filename: ", func(in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)

		files, _ := srv.ListFilesInWorkDir()
		for _, filename := range files {
			suggests = append(suggests, prompt.Suggest{Text: filename})
		}
		return prompt.FilterHasPrefix(suggests, in.Text, false)
	})

	return filename
}

func (srv *BufSrv) ListConflictedFilenames() ([]string, error) {
	workdirFilenames, err := srv.ListFilesInWorkDir()
	if err != nil {
		return make([]string, 0), err
	}
	bufDirFilenames, err := srv.ListFilesInBufDir()
	if err != nil {
		return make([]string, 0), err
	}

	duplicates := make([]string, 0)
	for _, filename := range bufDirFilenames {
		if slices.Contains(workdirFilenames, filename) {
			duplicates = append(duplicates, filename)
		}
	}

	return duplicates, nil
}
