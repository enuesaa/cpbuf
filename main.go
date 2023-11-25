package main

import (
	"log"

	"github.com/enuesaa/cpbuf/internal/cli"
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.NewRepos()
	log.SetFlags(0)

	app := &cobra.Command{
		Use:     "cpbuf",
		Short:   "A CLI tool to copy and paste files.\n`cpbuf` uses buf-dir to save files temporarily.",
		Version: "0.0.6",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	app.AddCommand(cli.CreateCopyCmd(repos))
	app.AddCommand(cli.CreateCCmd(repos))
	app.AddCommand(cli.CreatePasteCmd(repos))
	app.AddCommand(cli.CreatePCmd(repos))
	app.AddCommand(cli.CreateResetCmd(repos))
	app.AddCommand(cli.CreateListCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")
	app.SetHelpTemplate(`{{.Short}}
{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}
Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}
{{end}}
{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`)

	app.Execute()
}
