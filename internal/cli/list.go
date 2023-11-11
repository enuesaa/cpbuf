package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateListCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list filenames in buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			if !usecase.IsBufDirExist(repos) {
				return
			}

			files, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			for _, file := range files {
				fmt.Printf("%s\n", file.GetFilename())
			}
		},
	}

	return cmd
}
