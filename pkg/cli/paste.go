package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreatePasteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "paste",
		Short:   "Paste files into the current dir (alias: p)",
		Aliases: []string{"p"},
		RunE: func(cmd *cobra.Command, args []string) error {
			keep, _ := cmd.Flags().GetBool("keep")
			overwrite, _ := cmd.Flags().GetBool("overwrite")

			if !usecase.IsBufDirExist(repos) {
				fmt.Printf("No files were found.\n")
				return nil
			}
			filenames, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				return fmt.Errorf("failed to list files in buf dir.\n%s", err.Error())
			}
			if len(filenames) == 0 {
				fmt.Printf("No files were found.\n")
				if err := usecase.DeleteBufDir(repos); err != nil {
					return fmt.Errorf("failed to clear buf dir.\n%s", err.Error())
				}
				return nil
			}

			conflictedFilenames, err := usecase.ListConflictedFilenames(repos)
			if err != nil {
				return fmt.Errorf("failed to list files in current dir.\n%s", err.Error())
			}
			if len(conflictedFilenames) > 0 {
				if overwrite {
					for _, filename := range conflictedFilenames {
						if err := usecase.RemoveFileInWorkDir(repos, filename); err != nil {
							return fmt.Errorf("failed to remove a file in work dir.\n%s", err.Error())
						}
						fmt.Printf("removed: %s\n", filename)
					}
				} else {
					fmt.Print("These files already exist in this dir:\n")
					for _, filename := range conflictedFilenames {
						fmt.Printf("  - %s\n", filename)
					}
					fmt.Printf("\n")
					fmt.Printf("If you wish to overwrite them, run the following command:\n")
					fmt.Printf("  cpbuf paste --overwrite\n")
					return nil
				}
			}

			if err := usecase.Paste(repos); err != nil {
				return fmt.Errorf("failed to paste file.\n%s", err.Error())
			}
			if keep {
				return nil
			}
			if err := usecase.DeleteBufDir(repos); err != nil {
				return fmt.Errorf("failed to clear buf dir.\n%s", err.Error())
			}
			return nil
		},
	}
	cmd.Flags().Bool("keep", false, "keep buffered files. Default: false")
	cmd.Flags().BoolP("overwrite", "o", false, "overwrite conflicted files. Default: false")

	return cmd
}
