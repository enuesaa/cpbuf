package usecase

import (
	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/task"
)

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

func Buffer(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInWorkDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := registry.CopyToBufDir(file); err != nil {
			return err
		}
	}
	return nil
}

func ListFilesInBufDir(repos repository.Repos) ([]task.Bufferfile, error) {
	registry := task.NewRegistry(repos)
	return registry.ListFilesInBufDir()
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

func Paste(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInBufDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := registry.CopyToWorkDir(file); err != nil {
			return err
		}
	}
	return nil
}
