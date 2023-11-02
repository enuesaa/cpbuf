package repository

type FsMockRepository struct{}

func (repo *FsMockRepository) IsFileOrDirExist(path string) bool {
	return true
}

func (repo *FsMockRepository) Homedir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) Workdir() (string, error) {
	return "/", nil
}

func (repo *FsMockRepository) CreateDir(path string) error {
	return nil
}

func (repo *FsMockRepository) RemoveDir(path string) error {
	return nil
}

func (repo *FsMockRepository) CopyFile(srcPath string, dstPath string) error {
	return nil
}

func (repo *FsMockRepository) ListFiles(path string) ([]string, error) {
	return make([]string, 0), nil
}
