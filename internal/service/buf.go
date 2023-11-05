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
	return srv.repos.Fs.RemoveDir(path)
}

func (srv *BufSrv) CopyFileToBufDir(filename string) error {
	registryPath, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	dstPath := filepath.Join(registryPath, filename)
	return srv.repos.Fs.CopyFile(filename, dstPath)
}

func (srv *BufSrv) PasteFileToWorkDir(filename string) error {
	registryPath, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	filePathInBufDir := filepath.Join(registryPath, filename)
	return srv.repos.Fs.CopyFile(filePathInBufDir, filename)
}

func (srv *BufSrv) PasteFilesToWorkDir() error {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return err
	}
	filenames, err := srv.repos.Fs.ListFiles(bufDirPath)
	for _, filename := range filenames {
		if err := srv.PasteFileToWorkDir(filename); err != nil {
			return err
		}
	}
	return nil
}

func (srv *BufSrv) ListFilenames() ([]string, error) {
	bufDirPath, err := srv.GetBufDirPath()
	if err != nil {
		return make([]string, 0), err
	}
	return srv.repos.Fs.ListFiles(bufDirPath)
}

func (srv *BufSrv) ExtractSameFilenamesInWorkDir() ([]string, error) {
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

func (srv *BufSrv) ListFilesInWorkDir() ([]string, error) {
	workdirPath, err := srv.repos.Fs.Workdir()
	if err != nil {
		return make([]string, 0), err
	}

	return srv.repos.Fs.ListFiles(workdirPath)
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
