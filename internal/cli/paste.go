package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreatePasteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paste",
		Short: "paste file from tmp dir",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("paste files.")

			// paste all
			// if same filename exist, confirm
			// remove tmp dir
		},
	}

	return cmd
}
