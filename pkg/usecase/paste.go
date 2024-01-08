package usecase

import (
	"fmt"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/task"
)

func ListConflictedFilenames(repos repository.Repos) ([]string, error) {
	list := make([]string, 0)
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInBufDir()
	if err != nil {
		return list, err
	}
	for _, file := range files {
		if file.IsTopLevel() {
			exists, err := file.IsSameFileExistInWorkDir()
			if err != nil {
				return list, err
			}
			if exists {
				list = append(list, file.GetFilename())
			}
		}
	}
	return list, nil
}

func Paste(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInBufDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := registry.CopyToWorkDir(file); err != nil {
			return err
		}
		fmt.Printf("pasted: %s\n", file.GetFilename())
	}
	return nil
}
