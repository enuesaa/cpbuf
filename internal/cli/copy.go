package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/usecase"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "Copy file to buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive && len(args) == 0 {
				log.Printf("Error: please pass filename to copy.\n")
				return
			}
			if interactive {
				selected := usecase.SelectFileWithPrompt(repos)
				args = []string{selected}
			}
	
			if err := usecase.CreateBufDir(repos); err != nil {
				log.Printf("Error: failed to create buf dir.\n%s\n", err.Error())
				return
			}

			existFiles, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				log.Printf("Error: failed to list files in buf dir.\n")
				return
			}
			for _, file := range existFiles {
				fmt.Printf("buffered: %s\n", file.GetFilename())
			}

			filename := args[0]
			if err := usecase.Buffer(repos, filename); err != nil {
				log.Printf("Error: failed to copy files to buf dir.\n%s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "start interactive prompt and select file.")

	return cmd
}
