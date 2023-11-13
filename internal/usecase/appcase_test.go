package usecase

import (
	"testing"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestListConflictedFilenames(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/c",
			"/.cpbuf/d",
			"/.cpbuf/e",
			"/workdir/a",
			"/workdir/d",
			"/workdir/f",
		},
	}
	repos := repository.NewMockRepos(fsmock)

	actual, _ := ListConflictedFilenames(repos)
	assert.Equal(t, []string{"a", "d"}, actual)
}


func TestRemoveFileInWorkDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/c",
		},
	}
	repos := repository.NewMockRepos(fsmock)

	RemoveFileInWorkDir(repos, "a")
	files, _ := ListFilesInWorkDir(repos)
	actual := make([]string, 0)
	for _, file := range files {
		actual = append(actual, file.GetFilename())
	}
	assert.Equal(t, []string{"b", "c"}, actual)
}
