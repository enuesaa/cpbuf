package usecase

import (
	"fmt"
	"log"
	"strings"

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

func Buffer(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInWorkDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasPrefix(file.GetFilename(), filename) || filename == "." {
			if err := registry.CopyToBufDir(file); err != nil {
				if isBrokenSymlink, e := file.IsBrokenSymlink(); e != nil || !isBrokenSymlink {
					log.Printf("Error: %s\n", err)
					return err
				}
				fmt.Printf("WARNING: %s was ignored because this file seems to be a broken symlink.\n", file.GetFilename())
			}
			fmt.Printf("copied: %s\n", file.GetFilename())
		}
	}
	return nil
}