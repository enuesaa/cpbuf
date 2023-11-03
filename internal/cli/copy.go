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
			filename := ""
			if len(args) > 0 {
				filename = args[0]
			}
			if interactive {
				filename = repos.Fs.SelectFileWithPrompt()
			}
			bufSrv := service.NewBufSrv(repos)
			if err := bufSrv.CreateBufDir(); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			if err := bufSrv.CopyFileToBufDir(filename); err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "interactive")

	return cmd
}
