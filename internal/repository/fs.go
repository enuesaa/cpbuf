package repository

import (
	"io"
	"os"
	"path/filepath"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type FsRepositoryInterface interface {
	IsFileOrDirExist(path string) bool
	IsDir(path string) (bool, error)
	MkDir(path string) error
	Homedir() (string, error)
	Workdir() (string, error)
	StartSelectPrompt(message string, completer prompt.Completer) string
	CreateDir(path string) error
	Remove(path string) error
	CopyFile(srcPath string, dstPath string) error
	ListFiles(path string) ([]string, error)
	ListFilesRecursively(path string) ([]string, error)
}
type FsRepository struct{}

var termState *term.State

func (repo *FsRepository) IsFileOrDirExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsRepository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsRepository) MkDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
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

func (repo *FsRepository) Remove(path string) error {
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

func (repo *FsRepository) ListFilesRecursively(path string) ([]string, error) {
	filenames := make([]string, 0)
	err := filepath.Walk(path, func(fpath string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		filenames = append(filenames, fpath)
		return nil
	})

	return filenames, err
}

func (repo *FsRepository) StartSelectPrompt(message string, completer prompt.Completer) string {
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

	answer := prompt.Input(message, completer, options...)
	repo.restoreState()

	return answer
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
