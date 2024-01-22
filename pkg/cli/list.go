package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateListCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List files in buf dir",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !usecase.IsBufDirExist(repos) {
				return nil
			}

			files, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				return fmt.Errorf("failed to list files in buf dir.\n%s\n", err.Error())
			}

			for _, file := range files {
				fmt.Printf("%s\n", file.GetFilename())
			}
			return nil
		},
	}

	return cmd
}
