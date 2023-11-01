package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

type FshomeRepositoryInterface interface {
	IsRegistryExist(registryName string) bool
	CreateRegistry(registryName string) error
	IsFileExist(registryName string, path string) bool
	WriteFile(registryName string, path string, content string) error
	ReadFile(registryName string, path string) (string, error)
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

func (repo *FshomeRepository) IsFileExist(registryName string, path string) bool {
	if !repo.IsRegistryExist(registryName) {
		return false
	}
	homedir, err := repo.homedir()
	if err != nil {
		return false
	}
	fullpath := filepath.Join(homedir, registryName, path)
	return repo.isFileOrDirExist(fullpath)
}

func (repo *FshomeRepository) WriteFile(registryName string, path string, content string) error {
	if !repo.IsRegistryExist(registryName) {
		return fmt.Errorf("registry does not exist.")
	}
	homedir, err := repo.homedir()
	if err != nil {
		return err
	}
	fullpath := filepath.Join(homedir, registryName, path)
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte(content)); err != nil {
		return err
	}
	return nil
}

func (repo *FshomeRepository) ReadFile(registryName string, path string) (string, error) {
	if !repo.IsFileExist(registryName, path) {
		return "", fmt.Errorf("file does not exist.")
	}
	homedir, err := repo.homedir()
	if err != nil {
		return "", err
	}
	fullpath := filepath.Join(homedir, registryName, path)
	content, err := os.ReadFile(fullpath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
