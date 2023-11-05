package repository

import "github.com/c-bata/go-prompt"

type FsMockRepository struct {
	ListFilesInternal func(path string) []string
}

func (repo *FsMockRepository) IsFileOrDirExist(path string) bool {
	return true
}

func (repo *FsMockRepository) Homedir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) Workdir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	return nil
}

func (repo *FsMockRepository) RemoveDir(path string) error {
	return nil
}

func (repo *FsMockRepository) CopyFile(srcPath string, dstPath string) error {
	return nil
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	return repo.ListFilesInternal(path), nil
}

func (repo *FsMockRepository) StartSelectPrompt(message string, completer prompt.Completer) string {
	return ""
}
