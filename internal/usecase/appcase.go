package usecase

import (
	"fmt"
	"log"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/task"
)

func IsBufDirExist(repos repository.Repos) bool {
	registry := task.NewRegistry(repos)
	return registry.IsBufDirExist()
}

func CreateBufDir(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	if !registry.IsBufDirExist() {
		return registry.CreateBufDir()
	}
	return nil
}

func DeleteBufDir(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	if registry.IsBufDirExist() {
		return registry.DeleteBufDir()
	}
	return nil
}

func SelectFileWithPrompt(repos repository.Repos) string {
	registry := task.NewRegistry(repos)
	filename := repos.Prompt.StartSelectPrompt("filename: ", func(in prompt.Document) []prompt.Suggest {
		suggests := make([]prompt.Suggest, 0)

		files, err := registry.ListFilesInWorkDir()
		if err != nil {
			return suggests
		}
		for _, file := range files {
			suggests = append(suggests, prompt.Suggest{Text: file.GetFilename()})
		}
		return prompt.FilterHasPrefix(suggests, in.Text, false)
	})

	return filename
}

func Buffer(repos repository.Repos, filename string, force bool) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInWorkDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasPrefix(file.GetFilename(), filename) || filename == "." {
			if err := registry.CopyToBufDir(file); err != nil {
				log.Printf("Error: %s\n", err)
				if !force {
					return err
				}
			}
			fmt.Printf("copied: %s\n", file.GetFilename())
		}
	}
	return nil
}

func ListFilesInBufDir(repos repository.Repos) ([]task.Bufferfile, error) {
	registry := task.NewRegistry(repos)
	return registry.ListFilesInBufDir()
}

func ListFilesInWorkDir(repos repository.Repos) ([]task.Workfile, error) {
	registry := task.NewRegistry(repos)
	return registry.ListFilesInWorkDir()
}

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

func RemoveFileInWorkDir(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	return registry.RemoveFileInWorkDir(filename)
}

func RemoveFileInBufDir(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	return registry.RemoveFileInBufDir(filename)
}
