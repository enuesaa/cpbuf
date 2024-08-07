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
	repos := repository.NewMock(&fsmock)

	copyCmd := CreateCopyCmd(repos)
	copyCmd.SetArgs([]string{"a.txt"})
	copyCmd.Execute()
	// Output:
	// copied: a.txt
}

func ExampleCreateCopyCmd_multipleFiles() {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a.txt",
			"/workdir/b.txt",
			"/workdir/c.txt",
		},
	}
	repos := repository.NewMock(&fsmock)

	copyCmd := CreateCopyCmd(repos)
	copyCmd.SetArgs([]string{"."})
	copyCmd.Execute()
	// Output:
	// copied: a.txt
	// copied: b.txt
	// copied: c.txt
}

func ExampleCreateCopyCmd_useWildCard() {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a.txt",
			"/workdir/ab.txt",
			"/workdir/abc.txt",
			"/workdir/abcd.txt",
			"/workdir/bc.txt",
			"/workdir/ab",
			"/workdir/ab/aa.txt",
			"/workdir/abc",
			"/workdir/abc/aa.txt",
		},
	}
	repos := repository.NewMock(&fsmock)

	copyCmd := CreateCopyCmd(repos)
	copyCmd.SetArgs([]string{"ab*"})
	copyCmd.Execute()
	// Output:
	// copied: ab
	// copied: ab.txt
	// copied: ab/aa.txt
	// copied: abc
	// copied: abc.txt
	// copied: abc/aa.txt
	// copied: abcd.txt
}
