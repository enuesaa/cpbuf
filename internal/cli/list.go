package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/enuesaa/cpbuf/internal/service"
	"github.com/spf13/cobra"
)

func CreateListCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list filenames in buf dir",
		Run: func(cmd *cobra.Command, args []string) {
			bufSrv := service.NewBufSrv(repos)

			if !bufSrv.IsBufDirExist() {
				return
			}

			filenames, err := bufSrv.ListFilenames()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			for _, filename := range filenames {
				fmt.Printf("%s\n", filename)
			}
		},
	}

	return cmd
}
