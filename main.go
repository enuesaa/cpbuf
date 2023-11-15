package main

import (
	"github.com/enuesaa/cpbuf/internal/cli"
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.NewRepos()

	app := &cobra.Command{
		Use:     "cpbuf",
		Short:   "A CLI tool to copy and paste files.\n`cpbuf` uses buf-dir to save files temporarily.",
		Version: "0.0.5",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	app.AddCommand(cli.CreateCopyCmd(repos))
	app.AddCommand(cli.CreatePasteCmd(repos))
	app.AddCommand(cli.CreateClearCmd(repos))
	app.AddCommand(cli.CreateListCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.CompletionOptions.DisableDefaultCmd = true

	app.Execute()
}
