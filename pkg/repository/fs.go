package repository

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

type FsRepositoryInterface interface {
	IsExist(path string) bool
	IsDir(path string) (bool, error)
	GetModTime(path string) (time.Time, error)
	IsBrokenSymlink(path string) (bool, error)
	CreateDir(path string) error
	HomeDir() (string, error)
	WorkDir() (string, error)
	Remove(path string) error
	CopyFile(srcPath string, dstPath string) error
	ListFiles(path string) ([]string, error)
}
type FsRepository struct{}

func (repo *FsRepository) IsExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *FsRepository) IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return f.IsDir(), nil
}

func (repo *FsRepository) GetModTime(path string) (time.Time, error) {
	f, err := os.Stat(path)
	if err != nil {
		return time.Now(), err
	}
	return f.ModTime(), nil
}

func (repo *FsRepository) IsBrokenSymlink(path string) (bool, error) {
	f, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	// bitwise AND
	if f.Mode()&os.ModeSymlink != os.ModeSymlink {
		// Not a symlink
		return false, nil
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true, nil
	}
	return false, nil
}

func (repo *FsRepository) CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (repo *FsRepository) HomeDir() (string, error) {
	return os.UserHomeDir()
}

func (repo *FsRepository) WorkDir() (string, error) {
	return os.Getwd()
}

func (repo *FsRepository) Remove(path string) error {
	return os.RemoveAll(path)
}

func (repo *FsRepository) CopyFile(srcPath string, dstPath string) error {
	if err := repo.CreateDir(filepath.Dir(dstPath)); err != nil {
		return err
	}

	srcF, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	return err
}

func (repo *FsRepository) ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	filenames := make([]string, 0)
	for _, entry := range entries {
		if entry.Name() == ".git" {
			continue
		}
		filenames = append(filenames, filepath.Join(path, entry.Name()))
	}
	return filenames, nil
}
