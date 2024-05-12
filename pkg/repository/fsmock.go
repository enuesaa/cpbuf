package repository

import (
	"slices"
	"strings"
	"time"
)

type FsMockRepository struct {
	Files []string
}

//check file or dir existence with roughly logic.
//if path is dir, please include dir path inside FsMockRepository.Files.
func (repo *FsMockRepository) IsExist(path string) bool {
	return slices.Contains(repo.Files, path)
}

func (repo *FsMockRepository) IsDir(path string) (bool, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/workdir/" + path
	}
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			if filepath == path {
				continue
			}
			return true, nil
		}
	}
	return false, nil
}

func (repo *FsMockRepository) GetModTime(path string) (time.Time, error) {
	return time.Now(), nil
}

func (repo *FsMockRepository) IsBrokenSymlink(path string) (bool, error) {
	return false, nil
}

func (repo *FsMockRepository) HomeDir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) WorkDir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	repo.Files = append(repo.Files, path)
	return nil
}

func (repo *FsMockRepository) Remove(path string) error {
	list := make([]string, 0)
	for _, file := range repo.Files {
		if !strings.HasPrefix(file, path) {
			list = append(list, file)
		}
	}
	repo.Files = list
	return nil
}

func (repo *FsMockRepository) CopyFile(srcPath string, dstPath string) error {
	repo.Files = append(repo.Files, dstPath)
	return nil
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	list := make([]string, 0)
	for _, filepath := range repo.Files {
		if strings.HasPrefix(filepath, path) {
			rest := strings.TrimPrefix(filepath, path+"/")
			if !strings.Contains(rest, "/") {
				list = append(list, filepath)
			}
		}
	}

	return list, nil
}
