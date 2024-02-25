package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateResetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reset",
		Short:   "Clear buffered files (alias: r)",
		Aliases: []string{"r"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := usecase.DeleteBufDir(repos); err != nil {
				return fmt.Errorf("failed to clear buf dir.\n%s\n", err.Error())
			}
			return nil
		},
	}

	return cmd
}
