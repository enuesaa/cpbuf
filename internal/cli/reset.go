package cli

import (
	"log"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateResetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset [<filename>]",
		Short: "Clear buffered file. If filename is not passed, clear all files in buf dir.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				filename := args[0]
				if err := usecase.RemoveFileInBufDir(repos, filename); err != nil {
					log.Fatalf("Error: failed to remove a file in buf dir.\n%s\n", err.Error())
				}
				return
			}

			if err := usecase.DeleteBufDir(repos); err != nil {
				log.Fatalf("Error: failed to clear buf dir.\n%s\n", err.Error())
			}
		},
	}

	return cmd
}
