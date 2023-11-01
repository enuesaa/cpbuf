package repository

type Repos struct {
	Fshome   FshomeRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Fshome:   &FshomeRepository{},
	}
}
