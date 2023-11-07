package service

import (
	"path/filepath"
	"slices"

	"github.com/c-bata/go-prompt"
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

func (srv *BufSrv) GetBufDirPath() (string, error) {
	homedir, err := srv.repos.Fs.Homedir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homedir, ".cpbuf")
	return path, nil
}

func (srv *BufSrv) IsBufDirExist() bool {
	homedir, err := srv.repos.Fs.Homedir()
	if err != nil {
		return false
	}
	path := filepath.Join(homedir, ".cpbuf")
	return srv.repos.Fs.IsFileOrDirExist(path)
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

func (srv *BufSrv) CopyFileToBufDir(filename string) error {
	registryPath, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	dstPath := filepath.Join(registryPath, filename)
	return srv.copyFilesRecursively(filename, dstPath)
}

// todo: refactor
func (srv *BufSrv) copyFilesRecursively(path string, dstPath string) error {
	isDir, err := srv.repos.Fs.IsDir(path)
	if err != nil {
		return err
	}
	if isDir {
		files, err := srv.repos.Fs.ListFilesRecursively(path)
		if err != nil {
			return err
		}
		for _, file := range files {
			if err := srv.repos.Fs.MkDir(filepath.Dir(filepath.Join(filepath.Dir(dstPath), file))); err != nil {
				return err
			}
			isfiledir, err := srv.repos.Fs.IsDir(file)
			if err != nil {
				return err
			}
			if isfiledir {
				continue
			}
			if err := srv.repos.Fs.CopyFile(file, filepath.Join(filepath.Dir(dstPath), file)); err != nil {
				return err
			}
		}
		return nil
	}
	return srv.repos.Fs.CopyFile(path, dstPath)
}

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
		files, err := srv.repos.Fs.ListFilesRecursively(filePathInBufDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			relpath, err := filepath.Rel(registryPath, file)
			if err := srv.repos.Fs.MkDir(filepath.Dir(relpath)); err != nil {
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

func (srv *BufSrv) ListFilenames() ([]string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return make([]string, 0), err
	}
	return srv.repos.Fs.ListFiles(bufDirPath)
}

func (srv *BufSrv) ListConflictedFilenames() ([]string, error) {
	workdirPath, err := srv.repos.Fs.Workdir()
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

func (srv *BufSrv) RemoveFileInWorkDir(filename string) error {
	workdirPath, err := srv.repos.Fs.Workdir()
	if err != nil {
		return err
	}

	return srv.repos.Fs.Remove(filepath.Join(workdirPath, filename))
}

func (srv *BufSrv) ListFilesInWorkDir() ([]string, error) {
	workdirPath, err := srv.repos.Fs.Workdir()
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
	filename := srv.repos.Fs.StartSelectPrompt("filename: ", func (in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)
	
		files, _ := srv.repos.Fs.ListFiles(".")
		for _, filename := range files {
			suggests = append(suggests, prompt.Suggest{Text: filename})
		}
	
		return prompt.FilterHasPrefix(suggests, in.Text, false)
	})

	return filename
}
