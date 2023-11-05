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

			bufSrv := service.NewBufSrv(repos)
			if !bufSrv.IsBufDirExist() {
				fmt.Printf("(No files were found.)\n")
				return
			}
			alreadyExistsFilenames, err := bufSrv.ExtractSameFilenamesInWorkDir()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			if len(alreadyExistsFilenames) > 0 {
				fmt.Printf("These files are already exists in this dir.\n")
				for _, filename := range alreadyExistsFilenames {
					fmt.Printf("- %s\n", filename)
				}
				return
			}

			filenames, err := bufSrv.ListFilesInBufDir()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			if len(filenames) == 0 {
				fmt.Printf("(No files were found.)\n")
				// to delete buf dir, do not return here.
			}
			for _, filename := range filenames {
				if err := bufSrv.PasteFileToWorkDir(filename); err != nil {
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
	cmd.Flags().Bool("keep", false, "do not clear buf dir")

	return cmd
}
