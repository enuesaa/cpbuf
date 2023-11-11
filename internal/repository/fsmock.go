package repository

import (
	"fmt"
	"slices"
	"strings"
)

type FsMockRepository struct {
	Files []string
}

func (repo *FsMockRepository) IsExist(path string) bool {
	return slices.Contains(repo.Files, path)
}

func (repo *FsMockRepository) IsDir(path string) (bool, error) {
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			if filepath == path {
				// this is file. not dir.
				return false, nil
			}
			return true, nil
		}
	}

	return false, fmt.Errorf("file or dir does not exist.")
}

func (repo *FsMockRepository) HomeDir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) WorkDir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	return nil
}

func (repo *FsMockRepository) Remove(path string) error {
	return nil
}

func (repo *FsMockRepository) CopyFile(srcPath string, dstPath string) error {
	return nil
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	list := make([]string, 0)
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			rest := strings.TrimPrefix(filepath, path + "/")
			if !strings.Contains(rest, "/") {
				list = append(list, filepath)
			}
		}
	}

	return list, nil
}
