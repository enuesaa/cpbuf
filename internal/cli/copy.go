package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "copy file to tmp dir",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("copy files.")
		},
	}

	return cmd
}
