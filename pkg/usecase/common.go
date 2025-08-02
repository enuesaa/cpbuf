package usecase

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/task"
)

func IsBufDirExist(repos repository.Repos) bool {
	registry := task.NewRegistry(repos)
	return registry.IsBufDirExist()
}

func CreateBufDir(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	return registry.CreateBufDir()
}

func DeleteBufDir(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	return registry.DeleteBufDir()
}

func ListFilesInBufDir(repos repository.Repos) ([]task.Buffile, error) {
	registry := task.NewRegistry(repos)
	return registry.ListFilesInBufDir()
}

func HasFileInBufDir(repos repository.Repos) (bool, error) {
	registry := task.NewRegistry(repos)
	if !registry.IsBufDirExist() {
		return false, nil
	}
	list, err := registry.ListFilesInBufDir()
	if err != nil {
		return false, err
	}
	return len(list) > 0, nil
}

func RemoveFileInBufDir(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	return registry.RemoveFileInBufDir(filename)
}

func ListFilesInWorkDir(repos repository.Repos) ([]task.Workfile, error) {
	registry := task.NewRegistry(repos)
	return registry.ListFilesInWorkDir()
}

func RemoveFileInWorkDir(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	return registry.RemoveFileInWorkDir(filename)
}
