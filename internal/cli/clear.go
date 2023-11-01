package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/service"
	"github.com/spf13/cobra"
)

func CreateClearCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "clear buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			bufSrv := service.NewBufSrv(repos)

			if err := bufSrv.DeleteBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
			}
		},
	}

	return cmd
}
