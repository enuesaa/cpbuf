package usecase

import (
	"testing"

	"github.com/enuesaa/cpbuf/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestIsBufDirExist(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf",
		},
	}
	repos := repository.NewMock(&fsmock)
	assert.Equal(t, true, IsBufDirExist(repos))
}

func TestIsBufDirExist_NotExist(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{},
	}
	repos := repository.NewMock(&fsmock)
	assert.Equal(t, false, IsBufDirExist(repos))
}

func TestCreateBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{},
	}
	repos := repository.NewMock(&fsmock)
	assert.Nil(t, CreateBufDir(repos))
	assert.Equal(t, []string{"/.cpbuf"}, fsmock.Files)
}

func TestDeleteBufDir(t *testing.T) {
	fsmock := repository.FsMockRepository{
		Files: []string{
			"/.cpbuf",
			"/.cpbuf/a",
		},
	}
	repos := repository.NewMock(&fsmock)
	assert.Nil(t, DeleteBufDir(repos))
	assert.Equal(t, []string{}, fsmock.Files)
}
