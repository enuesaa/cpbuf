package repository

import (
	"io"
	"os"
	"path/filepath"
)

type FshomeRepositoryInterface interface {
	Workdir() (string, error)
	IsRegistryExist(registryName string) bool
	CreateRegistry(registryName string) error
	DeleteRegistry(registryName string) error
	GetResgistryPath(registryName string) (string, error)
	CopyFile(srcPath string, dstPath string) error
	ListFiles(path string) ([]string, error)
}
type FshomeRepository struct{}

func (repo *FshomeRepository) isFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FshomeRepository) homedir() (string, error) {
	return os.UserHomeDir()
}

func (repo *FshomeRepository) Workdir() (string, error) {
	return os.Getwd()
}

func (repo *FshomeRepository) GetResgistryPath(registryName string) (string, error) {
	homedir, err := repo.homedir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, registryName), nil
}

func (repo *FshomeRepository) IsRegistryExist(registryName string) bool {
	homedir, err := repo.homedir()
	if err != nil {
		return false
	}
	path := filepath.Join(homedir, registryName)
	return repo.isFileOrDirExist(path)
}

func (repo *FshomeRepository) CreateRegistry(registryName string) error {
	homedir, err := repo.homedir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, registryName)
	return os.Mkdir(path, 0755)
}

func (repo *FshomeRepository) DeleteRegistry(registryName string) error {
	homedir, err := repo.homedir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, registryName)
	return os.RemoveAll(path)
}

func (repo *FshomeRepository) CopyFile(srcPath string, dstPath string) error {
	srcF, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcF.Close()

    dstF, err := os.Create(dstPath)
    if err != nil {
        return err
    }
    defer dstF.Close()

    _, err = io.Copy(dstF, srcF)
	return err
}

func (repo *FshomeRepository) ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	filenames := make([]string, 0)
	for _, entry := range entries {
		filenames = append(filenames, entry.Name())
	}
	return filenames, nil
}

