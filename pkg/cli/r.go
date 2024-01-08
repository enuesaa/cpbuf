package cli

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateRCmd(repos repository.Repos) *cobra.Command {
	copyCmd := CreateResetCmd(repos)
	copyCmd.Use = "r"
	copyCmd.Short = "Alias for `reset`"

	return copyCmd
}
