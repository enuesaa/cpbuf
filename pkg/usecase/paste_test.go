package usecase

import (
	"testing"

	"github.com/enuesaa/cpbuf/pkg/repository"
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
	repos := repository.NewMock(&fsmock)

	actual, _ := ListConflictedFilenames(repos)
	assert.Equal(t, []string{"a", "d"}, actual)
}

func TestPaste(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/b",
			"/.cpbuf/c",
			"/workdir/a",
		},
	}
	repos := repository.NewMock(&fsmock)

	assert.Nil(t, Paste(repos))
	assert.Equal(t, fsmock.Files, []string{
		"/.cpbuf/b",
		"/.cpbuf/c",
		"/workdir/a",
		"/workdir/b",
		"/workdir/c",
	})
}
