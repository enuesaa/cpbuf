package usecase

import (
	"fmt"
	"slices"
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

func bufferFile(repos repository.Repos, filename string) error {
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
			if err := bufferFile(repos, file.GetFilename()); err != nil {
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
	dirs := splitRmEmpty(filename, "/")
	searchingDirs := splitRmEmpty(searching, "/")

	if len(dirs) == 0 || len(searchingDirs) == 0 {
		return len(dirs) == 0 && len(searchingDirs) == 0
	}

	for _, sd := range searchingDirs {
		dir := dirs[0]
		if !isTextMatch(dir, sd) {
			return false
		}
		if len(dirs) == 1 {
			break
		}
		dirs = dirs[1:]
	}
	return true
}

func isTextMatch(text string, pattern string) bool {
	if !strings.Contains(pattern, "*") {
		return text == pattern
	}

	pattern = strings.ReplaceAll(pattern, "*", "/*/")
	patternSplit := splitRmEmpty(pattern, "/")

	anythingOk := false
	for {
		if len(patternSplit) == 0 {
			break
		}
		ps := patternSplit[0]
		if ps == "*" {
			anythingOk = true
			patternSplit = slices.Delete(patternSplit, 0, 1)
			continue
		}
		var found bool
		text, found = strings.CutPrefix(text, ps)
		if found {
			anythingOk = false
			patternSplit = slices.Delete(patternSplit, 0, 1)
			continue
		}
		if !anythingOk {
			return false
		}
		if len(text) == 1 {
			text = ""
			break
		}
		text = text[1:]
	}

	if anythingOk {
		return len(patternSplit) == 0
	}
	return len(patternSplit) == 0 && len(text) == 0
}

func splitRmEmpty(s string, sep string) []string {
	split := strings.Split(s, sep)
	return slices.DeleteFunc(split, func(v string) bool {
		return v == ""
	})
}
