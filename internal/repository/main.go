package repository

type Repos struct {
	Fs FsRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Fs: &FsRepository{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		Fs: &FsMockRepository{
			ListFilesInternal: func(path string) []string {
				return make([]string, 0)
			},
		},
	}
}
