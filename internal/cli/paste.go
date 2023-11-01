package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreatePasteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paste",
		Short: "paste file from tmp dir",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("paste files.")
		},
	}

	return cmd
}
