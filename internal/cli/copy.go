package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/service"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy <filename>",
		Short: "copy file to buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			bufSrv := service.NewBufSrv(repos)
			if err := bufSrv.CreateBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			// cp <filename>
		},
	}
	cmd.Flags().BoolP("interactive", "-i", false, "interactive")

	return cmd
}
