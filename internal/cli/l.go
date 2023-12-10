package cli

import (
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreateLCmd(repos repository.Repos) *cobra.Command {
	copyCmd := CreateCopyCmd(repos)
	copyCmd.Use = "l"
	copyCmd.Short = "Alias for `list`"

	return copyCmd
}

