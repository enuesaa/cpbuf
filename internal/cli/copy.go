package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "copy file to buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("copy files.")

			// create tmp dir
			// cp <filename>
		},
	}
	cmd.Flags().BoolP("interactive", "-i", false, "interactive")

	return cmd
}
