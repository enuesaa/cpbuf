package cli

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
)

func ExampleCreateListCmd() {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf",
			"/.cpbuf/a.txt",
			"/.cpbuf/b.txt",
			"/.cpbuf/c.txt",
		},
	}
	repos := repository.NewMockRepos(&fsmock)

	listCmd := CreateListCmd(repos)
	listCmd.Execute()
	// Output:
	// a.txt
	// b.txt
	// c.txt
}
