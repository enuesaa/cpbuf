package cli

import (
	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/spf13/cobra"
)

func CreatePCmd(repos repository.Repos) *cobra.Command {
	copyCmd := CreatePasteCmd(repos)
	copyCmd.Use = "p"
	copyCmd.Short = "Alias for `paste`"

	return copyCmd
}
