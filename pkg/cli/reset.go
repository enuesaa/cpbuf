package cli

import (
	"log"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateResetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Clear buffered files",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.DeleteBufDir(repos); err != nil {
				log.Fatalf("Error: failed to clear buf dir.\n%s\n", err.Error())
			}
		},
	}

	return cmd
}
