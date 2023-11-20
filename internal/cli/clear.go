package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateResetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "clear buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.DeleteBufDir(repos); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}
		},
	}

	return cmd
}
