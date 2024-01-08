package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "Copy file to buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive && len(args) == 0 {
				log.Fatalf("Error: please pass filename to copy.\n")
			}
			if interactive {
				selected := usecase.SelectFileWithPrompt(repos)
				args = []string{selected}
			}
	
			if err := usecase.CreateBufDir(repos); err != nil {
				log.Fatalf("Error: failed to create buf dir.\n%s\n", err.Error())
			}

			existFiles, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				log.Fatalf("Error: failed to list files in buf dir.\n")
			}
			for _, file := range existFiles {
				fmt.Printf("buffered: %s\n", file.GetFilename())
			}

			for _, filename := range args {
				if err := usecase.Buffer(repos, filename); err != nil {
					log.Fatalf("Error: failed to copy files to buf dir.\n%s\n", err.Error())
				}
				fmt.Printf("copied: %s\n", filename)
			}
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "start interactive prompt and select file.")

	return cmd
}
