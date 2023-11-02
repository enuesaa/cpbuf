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
			bufSrv := service.NewBufSrv(repos)

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

			if err := bufSrv.PasteFilesToWorkDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			if err := bufSrv.DeleteBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}
		},
	}
	// cmd.Flags().Bool("keep", false, "do not clear buf dir")

	return cmd
}
