package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/service"
	"github.com/spf13/cobra"
)

func CreatePasteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paste",
		Short: "paste files",
		Run: func(cmd *cobra.Command, args []string) {
			keep, _ := cmd.Flags().GetBool("keep")
			overwrite, _ := cmd.Flags().GetBool("overwrite")

			bufSrv := service.NewBufSrv(repos)
			if !bufSrv.IsBufDirExist() {
				fmt.Printf("No files were found.\n")
				return
			}

			conflictedFilenames, err := bufSrv.ListConflictedFilenames()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			if len(conflictedFilenames) > 0 {
				if overwrite {
					for _, filename := range conflictedFilenames {
						if err := bufSrv.RemoveFileInWorkDir(filename); err != nil {
							fmt.Printf("error: %s\n", err.Error())
							return
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

			filenames, err := bufSrv.ListFilesInBufDir()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			if len(filenames) == 0 {
				fmt.Printf("No files were found.\n")
				// to delete buf dir, do not return here.
			}
			for _, filename := range filenames {
				if err := bufSrv.Paste(filename); err != nil {
					fmt.Printf("error: %s\n", err.Error())
					return
				}
				fmt.Printf("pasted: %s\n", filename)
			}

			if keep {
				return
			}

			if err := bufSrv.DeleteBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}
		},
	}
	cmd.Flags().Bool("keep", false, "keep buffered files. Default: false")
	cmd.Flags().BoolP("overwrite", "o", false, "overwrite conflicted files. Default: false")

	return cmd
}
