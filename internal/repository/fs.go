package repository

import (
	"io"
	"os"
	"fmt"
	// "path/filepath"
	// "strings"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type FsRepositoryInterface interface {
	IsFileOrDirExist(path string) bool
	Homedir() (string, error)
	Workdir() (string, error)
	SelectFileWithPrompt() string
	CreateDir(path string) error
	RemoveDir(path string) error
	CopyFile(srcPath string, dstPath string) error
	ListFiles(path string) ([]string, error)
}
type FsRepository struct{}

var termState *term.State

func (repo *FsRepository) IsFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsRepository) Homedir() (string, error) {
	return os.UserHomeDir()
}

func (repo *FsRepository) Workdir() (string, error) {
	return os.Getwd()
}

func (repo *FsRepository) CreateDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

func (repo *FsRepository) RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func (repo *FsRepository) CopyFile(srcPath string, dstPath string) error {
	srcF, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	return err
}

func (repo *FsRepository) ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	filenames := make([]string, 0)
	for _, entry := range entries {
		filenames = append(filenames, entry.Name())
	}
	return filenames, nil
}

// see https://github.com/c-bata/go-prompt/issues/8
// see https://github.com/c-bata/go-prompt/issues/233
func (repo *FsRepository) saveState() {
	state, _ := term.GetState(int(os.Stdin.Fd()))
	termState = state
}

func (repo *FsRepository) restoreState() {
	if termState != nil {
		term.Restore(int(os.Stdin.Fd()), termState)
	}
	termState = nil
}

func (repo *FsRepository) SelectFileWithPrompt() string {
	repo.saveState()

	options := make([]prompt.Option, 0)
	options = append(options, prompt.OptionAddKeyBind(prompt.KeyBind{
		Key: prompt.ControlC,
		Fn: func(*prompt.Buffer) {
			repo.restoreState()
			os.Exit(0)
		},
	}))
	options = append(options, prompt.OptionShowCompletionAtStart())
	options = append(options, prompt.OptionSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionScrollbarThumbColor(prompt.Black))
	options = append(options, prompt.OptionSuggestionTextColor(prompt.White))
	options = append(options, prompt.OptionSelectedSuggestionBGColor(prompt.Black))
	options = append(options, prompt.OptionSelectedSuggestionTextColor(prompt.Cyan))
	options = append(options, prompt.OptionMaxSuggestion(15))
	options = append(options, prompt.OptionPrefixTextColor(prompt.Brown))
	options = append(options, prompt.OptionCompletionOnDown())

	for {
		dir := prompt.Input("filename: ", repo.suggestDirs, options...)
		if repo.IsFileOrDirExist(dir) {
			repo.restoreState()
			return dir
		}
		fmt.Printf("Dir %s does not exist. \n", dir)
	}
}

func (repo *FsRepository) suggestDirs(in prompt.Document) []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)

	files, _ := repo.ListFiles(".")
	for _, filename := range files {
		suggests = append(suggests, prompt.Suggest{ Text: filename })
	}

	return prompt.FilterHasPrefix(suggests, in.Text, false)
}
