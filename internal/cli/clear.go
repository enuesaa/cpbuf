package cli

import (
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreateClearCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "clear buf dir",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
