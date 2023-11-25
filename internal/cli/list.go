package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateListCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List files in buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			if !usecase.IsBufDirExist(repos) {
				return
			}

			files, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				log.Printf("Error: failed to list files in buf dir.\n%s\n", err.Error())
				return
			}

			for _, file := range files {
				fmt.Printf("%s\n", file.GetFilename())
			}
		},
	}

	return cmd
}
