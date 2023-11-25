package cli

import (
	"log"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Aliases: []string{"c"},
		Short: "Copy file to buf dir. Alias: c",
		Run: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive && len(args) == 0 {
				log.Printf("Error: please pass filename to copy.\n")
				return
			}

			if interactive {
				selected := usecase.SelectFileWithPrompt(repos)
				args = []string{selected}
			}

			if err := usecase.CreateBufDir(repos); err != nil {
				log.Printf("Error: failed to create buf dir.\n%s\n", err.Error())
				return
			}

			filename := args[0]
			if err := usecase.Buffer(repos, filename); err != nil {
				log.Printf("Error: failed to copy files to buf dir.\n%s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "start interactive prompt and select file.")

	return cmd
}
