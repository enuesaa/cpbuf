package cli

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
)

func ExampleCreateCopyCmd() {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a.txt",
		},
	}
	repos := repository.NewMockRepos(&fsmock)

	copyCmd := CreateCopyCmd(repos)
	copyCmd.SetArgs([]string{"a.txt"})
	copyCmd.Execute()
	// Output:
	// copied: a.txt
}
