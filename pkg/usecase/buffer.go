package usecase

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/task"
)

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

func BufferFile(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	isBrokenSymlink, err := registry.IsBrokenSymlink(filename)
	if err != nil {
		return err
	}
	if isBrokenSymlink {
		fmt.Printf("WARNING: %s was ignored because this file seems to be a broken symlink.\n", filename)
		return nil
	}
	file, err := registry.GetWorkfileWithFilename(filename)
	if err != nil {
		return err
	}
	if err := registry.CopyToBufDir(file); err != nil {
		return err
	}
	fmt.Printf("copied: %s\n", file.GetFilename())
	return nil
}

func BufferAll(repos repository.Repos) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInWorkDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := BufferFile(repos, file.GetFilename()); err != nil {
			return err
		}
	}
	return nil
}
