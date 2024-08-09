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
	repos := repository.NewMock(&fsmock)

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
	repos := repository.NewMock(&fsmock)

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

func TestIsSearchingFile(t *testing.T) {
	assert.Equal(t, true, isSearchingFile("a.txt", "a.txt"))
	assert.Equal(t, true, isSearchingFile("a.txt", "."))
	assert.Equal(t, true, isSearchingFile("a.txt", "*"))
	assert.Equal(t, true, isSearchingFile("a.txt", "a*"))
	assert.Equal(t, true, isSearchingFile("a.txt", "*.txt"))
	assert.Equal(t, true, isSearchingFile("a/bb.txt", "a"))
	assert.Equal(t, true, isSearchingFile("a/bb/cc.txt", "a"))
	assert.Equal(t, true, isSearchingFile("a/bb/cc.txt", "a*"))
	assert.Equal(t, true, isSearchingFile("ab/bb/cc.txt", "a*"))
	assert.Equal(t, true, isSearchingFile("ab/bb/cc.txt", "ab"))
	assert.Equal(t, true, isSearchingFile("ab/bb/cc.txt", "ab/bb"))
	assert.Equal(t, true, isSearchingFile("ab/bb/cc.txt", "ab/bb/cc.txt"))

	assert.Equal(t, false, isSearchingFile("a.txt", "b.txt"))
	assert.Equal(t, false, isSearchingFile("a.txt", ""))
	assert.Equal(t, false, isSearchingFile("a.txt", "b*"))
	assert.Equal(t, false, isSearchingFile("a.txt", "a.txta"))
	assert.Equal(t, false, isSearchingFile("ab/bb/cc.txt", "a"))
	assert.Equal(t, false, isSearchingFile("ab/bb/cc.txt", "*a"))
}

func TestIsTextMatch(t *testing.T) {
	assert.Equal(t, true, isTextMatch("a.txt", "a.txt"))
	assert.Equal(t, true, isTextMatch("a.txt", "*.txt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a.*xt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a*xt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a*"))

	assert.Equal(t, false, isTextMatch("ab", "a"))
	assert.Equal(t, false, isTextMatch("ab", "*a"))
	assert.Equal(t, false, isTextMatch("a.txt", "b*"))
	assert.Equal(t, false, isTextMatch("a.txt", "*c"))
}
