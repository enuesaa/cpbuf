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
			
			// paste all
			// if same filename exist, confirm
			
			if err := bufSrv.DeleteBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}

			// remove tmp dir
		},
	}
	cmd.Flags().Bool("keep", false, "do not clear buf dir")

	return cmd
}
