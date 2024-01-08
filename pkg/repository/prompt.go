package repository

import (
	"os"

	"github.com/c-bata/go-prompt"
	"golang.org/x/term"
)

type PromptRepositoryInterface interface {
	StartSelectPrompt(message string, completer prompt.Completer) string
}
type PromptRepository struct {
	termState *term.State
}

func (repo *PromptRepository) StartSelectPrompt(message string, completer prompt.Completer) string {
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
func (repo *PromptRepository) saveState() {
	state, _ := term.GetState(int(os.Stdin.Fd()))
	repo.termState = state
}

func (repo *PromptRepository) restoreState() {
	if repo.termState != nil {
		term.Restore(int(os.Stdin.Fd()), repo.termState)
	}
	repo.termState = nil
}
