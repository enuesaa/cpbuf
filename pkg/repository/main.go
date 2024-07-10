package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptRepositoryInterface
}

func New() Repos {
	return Repos{
		Fs:     &FsRepository{},
		Prompt: &PromptRepository{},
	}
}

func NewMock(fsmock *FsMockRepository) Repos {
	return Repos{
		Fs:     fsmock,
		Prompt: &PromptRepository{},
	}
}
