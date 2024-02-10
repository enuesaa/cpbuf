package cli

import (
	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/enuesaa/cpbuf/pkg/usecase"
)

func ExampleCreateCopyCmd() {
	usecase.DeleteBufDir(repository.NewRepos())

	copyCmd := CreateCopyCmd(repository.NewRepos())
	copyCmd.SetArgs([]string{"c.go"})
	copyCmd.Execute()
	// Output:
	// copied: c.go
}
