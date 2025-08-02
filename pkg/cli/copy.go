package cli

import (
	"fmt"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateCopyCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "copy <filename>",
		Short:   "Copy files to the buf dir (alias: c)",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive && len(args) == 0 {
				return fmt.Errorf("please pass filename to copy")
			}
			if interactive {
				selected := usecase.SelectFileWithPrompt(repos)
				args = []string{selected}
			}

			if err := usecase.CreateBufDir(repos); err != nil {
				return fmt.Errorf("failed to create buf dir.\n%s", err.Error())
			}

			existingFiles, err := usecase.ListFilesInBufDir(repos)
			if err != nil {
				return fmt.Errorf("failed to list files in buf dir")
			}
			for _, filename := range args {
				if err := usecase.Buffer(repos, filename); err != nil {
					return fmt.Errorf("failed to copy files to buf dir.\n%s", err.Error())
				}
			}

			if len(existingFiles) == 0 {
				return nil
			}
			fmt.Printf("\n")
			fmt.Printf("WARNING: These files already copied.\n")
			for _, file := range existingFiles {
				fmt.Printf("* copied on %s: %s\n", file.GetBufferedDate(), file.GetFilename())
			}
			return nil
		},
	}
	cmd.Flags().BoolP("interactive", "i", false, "start interactive prompt and select file.")

	return cmd
}
