package cli

import (
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreateCCmd(repos repository.Repos) *cobra.Command {
	copyCmd := CreateCopyCmd(repos)
	copyCmd.Use = "c <filename>"
	copyCmd.Short = "Alias for `copy`"

	return copyCmd
}
