package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreatePasteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paste",
		Short: "Paste files to current dir",
		Run: func(cmd *cobra.Command, args []string) {
			keep, _ := cmd.Flags().GetBool("keep")
			overwrite, _ := cmd.Flags().GetBool("overwrite")

			if !usecase.IsBufDirExist(repos) {
				fmt.Printf("No files were found.\n")
				return
			}

			conflictedFilenames, err := usecase.ListConflictedFilenames(repos)
			if err != nil {
				log.Fatalf("Error: failed to list files in current dir.\n%s\n", err.Error())
			}
			if len(conflictedFilenames) > 0 {
				if overwrite {
					for _, filename := range conflictedFilenames {
						if err := usecase.RemoveFileInWorkDir(repos, filename); err != nil {
							log.Fatalf("Error: failed to remove a file in work dir.\n%s\n", err.Error())
						}
						fmt.Printf("removed: %s\n", filename)
					}
				} else {
					fmt.Printf("These files already exist in this dir.\n")
					for _, filename := range conflictedFilenames {
						fmt.Printf("- %s\n", filename)
					}
					fmt.Printf("\n")
					fmt.Printf("If you wish overwrite these files, please run command below.\n")
					fmt.Printf("  cpbuf paste --overwrite\n")
					return
				}
			}

			filenames, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				log.Fatalf("Error: failed to list files in buf dir.\n%s\n", err.Error())
			}
			if len(filenames) == 0 {
				fmt.Printf("No files were found.\n")
				// to delete buf dir, do not return here.
			} else {
				if err := usecase.Paste(repos); err != nil {
					log.Fatalf("Error: failed to paste file.\n%s\n", err.Error())
				}
			}

			if keep {
				return
			}

			if err := usecase.DeleteBufDir(repos); err != nil {
				log.Fatalf("Error: failed to clear buf dir.\n%s\n", err.Error())
			}
		},
	}
	cmd.Flags().Bool("keep", false, "keep buffered files. Default: false")
	cmd.Flags().BoolP("overwrite", "o", false, "overwrite conflicted files. Default: false")

	return cmd
}
