package service

import (
	"testing"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/stretchr/testify/assert"
)

// see https://qiita.com/takehanKosuke/items/4342ca544d205fb36eb0

func TestGetBufDirPath(t *testing.T) {
	fsmock := repository.FsMockRepository{ Files: []string{} }
	repos := repository.NewMockRepos(fsmock)
	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.GetBufDirPath()
	assert.Equal(t, "/.cpbuf", actual)
}

func TestIsBufDirExist(t *testing.T) {
	fsmock := repository.FsMockRepository{ Files: []string{"/.cpbuf"} }
	repos := repository.NewMockRepos(fsmock)
	bufSrv := NewBufSrv(repos)
	assert.Equal(t, true, bufSrv.IsBufDirExist())
}

func TestCreateBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{ Files: []string{} }
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	bufSrv.CreateBufDir()
	assert.Equal(t, true, bufSrv.IsBufDirExist())
}

func TestDeleteBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{ Files: []string{"/.cpbuf"} }
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	bufSrv.DeleteBufDir()
	assert.Equal(t, false, bufSrv.IsBufDirExist())
}

func TestBufferFile(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/b/bb",
			"/workdir/b/bb/bbb",
		},
	}
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	assert.Equal(t, nil, bufSrv.Buffer("a"))
	assert.Equal(t, nil, bufSrv.Buffer("b"))

	actual, _ := bufSrv.ListFilesRecursively("/.cpbuf")
	assert.Equal(t, []string{"/.cpbuf/a", "/.cpbuf/b", "/.cpbuf/b/bb", "/.cpbuf/b/bb/bbb"}, actual)
}

func TestListFilesRecursively(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/b/bb",
			"/.cpbuf/b/bb/bbb",
			"/workdir/c",
		},
	}
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListFilesRecursively("/.cpbuf")
	assert.Equal(t, []string{"/.cpbuf/a", "/.cpbuf/b", "/.cpbuf/b/bb", "/.cpbuf/b/bb/bbb"}, actual)
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

	bufSrv := NewBufSrv(repos)
	bufSrv.RemoveFileInWorkDir("a")
	actual, _ := bufSrv.ListFilesInWorkDir()
	assert.Equal(t, []string{"b", "c"}, actual)
}

func TestListFilesInWorkDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/c",
		},
	}
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListFilesInWorkDir()
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
	repos := repository.NewMockRepos(fsmock)

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
	repos := repository.NewMockRepos(fsmock)

	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.ListConflictedFilenames()
	assert.Equal(t, []string{"a", "d"}, actual)
}
