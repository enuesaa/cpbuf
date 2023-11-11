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
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive && len(args) == 0 {
				fmt.Printf("error: please pass filename to copy.\n")
				return
			}

			bufSrv := service.NewBufSrv(repos)
			if interactive {
				selected := bufSrv.SelectFileWithPrompt()
				args = []string{selected}
			}

			if !bufSrv.IsBufDirExist() {
				if err := bufSrv.CreateBufDir(); err != nil {
					fmt.Printf("error: %s\n", err.Error())
					return
				}
			}

			filename := args[0]
			if err := bufSrv.Buffer(filename); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "start interactive prompt and select file.")

	return cmd
}
