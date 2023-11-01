package main

import (
	"github.com/enuesaa/cpbuf/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "cpbuf",
		Short:   "A CLI tool to copy and paste files",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	app.AddCommand(cli.CreateCopyCmd())

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.CompletionOptions.DisableDefaultCmd = true

	app.Execute()
}
