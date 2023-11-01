package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "copy file to tmp dir",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("copy files.")
		},
	}

	return cmd
}
