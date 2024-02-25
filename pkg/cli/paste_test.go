package cli

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
)

func ExampleCreatePasteCmd() {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf",
			"/.cpbuf/a.txt",
			"/workdir/b.txt",
		},
	}
	repos := repository.NewMockRepos(&fsmock)

	pasteCmd := CreatePasteCmd(repos)
	pasteCmd.Execute()
	// Output:
	// pasted: a.txt
}
