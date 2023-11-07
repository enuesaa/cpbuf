package repository

type FsMockRepository struct {
	ListFilesInternal func(path string) []string
}

func (repo *FsMockRepository) IsFileOrDirExist(path string) bool {
	return true
}

func (repo *FsMockRepository) IsDir(path string) (bool, error) {
	return false, nil
}

func (repo *FsMockRepository) Homedir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) Workdir() (string, error) {
	return "/workdir", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	return nil
}

func (repo *FsMockRepository) Remove(path string) error {
	return nil
}

func (repo *FsMockRepository) CopyFile(srcPath string, dstPath string) error {
	return nil
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	return repo.ListFilesInternal(path), nil
}

func (repo *FsMockRepository) ListFilesRecursively(path string) ([]string, error) {
	return make([]string, 0), nil
}
