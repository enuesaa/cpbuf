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

func NewMockRepos(files []string) Repos {
	return Repos{
		Fs: &FsMockRepository{
			Files: files,
		},
		Prompt: &PromptRepository{},
	}
}
