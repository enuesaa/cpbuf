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

func NewMockRepos(fsmock FsMockRepository) Repos {
	return Repos{
		Fs: &fsmock,
		Prompt: &PromptRepository{},
	}
}
