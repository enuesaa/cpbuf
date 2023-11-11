package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:     &FsRepository{},
		Prompt: &PromptRepository{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		Fs: &FsMockRepository{
			// ListFilesInternal: func(path string) []string {
			// 	return make([]string, 0)
			// },
			Files: make([]string, 0),
		},
		Prompt: &PromptRepository{},
	}
}
