package service

import (
	"testing"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetBufDirPath(t *testing.T) {
	repos := repository.NewMockRepos([]string{})
	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.GetBufDirPath()
	assert.Equal(t, "/.cpbuf", actual)
}

func TestIsBufDirExist(t *testing.T) {
	repos := repository.NewMockRepos([]string{"/.cpbuf"})
	bufSrv := NewBufSrv(repos)
	assert.Equal(t, true, bufSrv.IsBufDirExist())
}

func ListFilesInWorkDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/c",
		},
	}
	repos := repository.Repos{
		Fs: &fsmock,
	}

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListFilesInBufDir()
	assert.Equal(t, []string{"a", "b", "c"}, actual)
}

func TestListFilesInBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/c",
		},
	}
	repos := repository.Repos{
		Fs: &fsmock,
	}

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListFilesInBufDir()
	assert.Equal(t, []string{"a", "b", "c"}, actual)
}

func TestExtractSameFilenamesInWorkDir(t *testing.T) {
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
	repos := repository.Repos{
		Fs: &fsmock,
	}

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListConflictedFilenames()
	assert.Equal(t, []string{"a", "d"}, actual)
}
