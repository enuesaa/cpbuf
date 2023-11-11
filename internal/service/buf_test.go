package service

import (
	"testing"

	"github.com/enuesaa/cpbuf/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetBufDirPath(t *testing.T) {
	repos := repository.NewMockRepos()
	bufSrv := NewBufSrv(repos)
	actual, _ := bufSrv.GetBufDirPath()
	assert.Equal(t, actual, "/.cpbuf")
}

func TestIsBufDirExist(t *testing.T) {
	repos := repository.NewMockRepos()
	bufSrv := NewBufSrv(repos)
	assert.Equal(t, bufSrv.IsBufDirExist(), true)
}

// func TestExtractSameFilenamesInWorkDir(t *testing.T) {
// 	repos := repository.Repos{
// 		Fs: &repository.FsMockRepository{
// 			ListFilesInternal: func(path string) []string {
// 				if path == "/.cpbuf" {
// 					return []string{"a", "b", "c", "d", "e"}
// 				}
// 				if path == "/workdir" {
// 					return []string{"a", "d", "f"}
// 				}
// 				return make([]string, 0)
// 			},
// 		},
// 	}

// 	bufSrv := NewBufSrv(repos)
// 	actual, _ := bufSrv.ListConflictedFilenames()
// 	assert.Equal(t, actual, []string{"a", "d"})
// }
