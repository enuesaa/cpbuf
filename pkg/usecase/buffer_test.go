package usecase

import (
	"testing"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestBufferFile(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
		},
	}
	repos := repository.NewMockRepos(&fsmock)

	assert.Nil(t, Buffer(repos, "a"))
	assert.Equal(t, fsmock.Files, []string{
		"/workdir/a",
		"/.cpbuf/a",
	})
}

func TestBuffer(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/a/b",
			"/workdir/a/c",
		},
	}
	repos := repository.NewMockRepos(&fsmock)

	assert.Nil(t, Buffer(repos, "a"))
	assert.Equal(t, fsmock.Files, []string{
		"/workdir/a",
		"/workdir/a/b",
		"/workdir/a/c",
		"/.cpbuf/a",
		"/.cpbuf/a/b",
		"/.cpbuf/a/c",
	})
}
