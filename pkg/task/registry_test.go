package task

import (
	"testing"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetBufDirPath(t *testing.T) {
	fsmock := repository.FsMockRepository{Files: []string{}}
	repos := repository.NewMock(&fsmock)
	registry := NewRegistry(repos)
	actual, _ := registry.GetBufDirPath()
	assert.Equal(t, "/.cpbuf", actual)
}

func TestIsBufDirExist(t *testing.T) {
	fsmock := repository.FsMockRepository{Files: []string{"/.cpbuf"}}
	repos := repository.NewMock(&fsmock)
	registry := NewRegistry(repos)
	assert.Equal(t, true, registry.IsBufDirExist())
}

func TestCreateBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{Files: []string{}}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.CreateBufDir()
	assert.Equal(t, true, registry.IsBufDirExist())
}

func TestDeleteBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{Files: []string{"/.cpbuf"}}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.DeleteBufDir()
	assert.Equal(t, false, registry.IsBufDirExist())
}

func TestListFilesInBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/c",
		},
	}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	buffiles, _ := registry.ListFilesInBufDir()
	actual := make([]string, 0)
	for _, buffile := range buffiles {
		actual = append(actual, buffile.GetFilename())
	}
	assert.Equal(t, []string{"a", "b", "c"}, actual)
}

func TestListFilesInWorkDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/c",
		},
	}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	workfiles, _ := registry.ListFilesInWorkDir()
	actual := make([]string, 0)
	for _, workfile := range workfiles {
		actual = append(actual, workfile.GetFilename())
	}
	assert.Equal(t, []string{"a", "b", "c"}, actual)
}

func TestCopyToBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/workdir/a",
			"/workdir/b",
			"/workdir/b/bb",
			"/workdir/b/bb/bbb",
			"/workdir/c",
		},
	}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.CopyToBufDir(NewWorkfile(repos, "/workdir/a", "/.cpbuf", "/workdir"))
	registry.CopyToBufDir(NewWorkfile(repos, "/workdir/b", "/.cpbuf", "/workdir"))

	actual, _ := registry.ListFilesRecursively("/.cpbuf")
	assert.Equal(t, []string{"/.cpbuf/a", "/.cpbuf/b"}, actual)
}

func TestCopyToWorkDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/b/bb",
			"/.cpbuf/b/bb/bbb",
			"/.cpbuf/c",
		},
	}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.CopyToWorkDir(NewBufferfile(repos, "/.cpbuf/a", "/.cpbuf", "/workdir"))
	registry.CopyToWorkDir(NewBufferfile(repos, "/.cpbuf/b", "/.cpbuf", "/workdir"))

	actual, _ := registry.ListFilesRecursively("/workdir")
	assert.Equal(t, []string{"/workdir/a", "/workdir/b"}, actual)
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
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	actual, _ := registry.ListFilesRecursively("/.cpbuf")
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
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.RemoveFileInWorkDir("a")
	files, _ := registry.ListFilesInWorkDir()
	actual := make([]string, 0)
	for _, file := range files {
		actual = append(actual, file.GetFilename())
	}
	assert.Equal(t, []string{"b", "c"}, actual)
}

func TestRemoveFileInBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf/a",
			"/.cpbuf/b",
			"/.cpbuf/c",
		},
	}
	repos := repository.NewMock(&fsmock)

	registry := NewRegistry(repos)
	registry.RemoveFileInBufDir("a")
	files, _ := registry.ListFilesInBufDir()
	actual := make([]string, 0)
	for _, file := range files {
		actual = append(actual, file.GetFilename())
	}
	assert.Equal(t, []string{"b", "c"}, actual)
}
