package usecase

import (
	"fmt"
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

func BufferFile(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	file := registry.GetWorkfile(filename)
	if err := file.CheckExist(); err != nil {
		if is, e := file.IsBrokenSymlink(); is && e == nil {
			fmt.Printf("WARNING: %s was ignored because this file seems to be a broken symlink.\n", filename)
			return nil
		}
		return err
	}

	if err := registry.CopyToBufDir(file); err != nil {
		return err
	}
	fmt.Printf("copied: %s\n", file.GetFilename())
	return nil
}

// Copy a file to buf dir.
//
// filename accepts multiple format below.
// - aa.txt: copy aa.txt
// - .     : copy all files in current dir
// - *     : copy all files in current dir
// - *a.txt: copy all files which ends with a.txt in current dir
func Buffer(repos repository.Repos, filename string) error {
	registry := task.NewRegistry(repos)
	files, err := registry.ListFilesInWorkDir()
	if err != nil {
		return err
	}
	for _, file := range files {
		if isSearchingFile(file.GetFilename(), filename) {
			if err := BufferFile(repos, file.GetFilename()); err != nil {
				return err
			}	
		}
	}
	return nil
}

func isSearchingFile(filename string, searching string) bool {
	if searching == "." || searching == "*" {
		return true
	}
	if strings.Contains(filename, "/") {
		return strings.HasPrefix(filename, searching)
	}
	if strings.Contains(searching, "*") {
		filenameSplit := strings.Split(filename, "") // like a.txt
		searchingSplit := strings.Split(strings.ReplaceAll(searching, "*", ""), "") // like a* or *txt
		for _, char := range filenameSplit {
			if searchingSplit[0] == char {
				if len(searchingSplit) > 1 {
					searchingSplit = searchingSplit[1:]
				} else {
					return true
				}
			}
		}
		return len(searchingSplit) == 0
	}

	return filename == searching
}
