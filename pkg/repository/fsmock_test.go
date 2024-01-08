package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	cases := []struct {
		files  []string
		check  string
		expect bool
	}{
		{
			files:  []string{"/aaa"},
			check:  "/aaa",
			expect: true,
		},
		{
			files:  []string{"/bbb"},
			check:  "/bbb",
			expect: true,
		},
	}

	for _, tt := range cases {
		fsmock := FsMockRepository{
			Files: tt.files,
		}
		assert.Equal(t, tt.expect, fsmock.IsExist(tt.check))
	}
}

func TestIsDir(t *testing.T) {
	cases := []struct {
		files  []string
		check  string
		expect bool
	}{
		{
			files:  []string{"/aaa", "/aaa/bb"},
			check:  "/aaa",
			expect: true,
		},
		{
			files:  []string{"/bbb"},
			check:  "/bbb",
			expect: false,
		},
	}

	for _, tt := range cases {
		fsmock := FsMockRepository{
			Files: tt.files,
		}
		actual, _ := fsmock.IsDir(tt.check)
		assert.Equal(t, tt.expect, actual)
	}
}

func TestListFiles(t *testing.T) {
	cases := []struct {
		files  []string
		check  string
		expect []string
	}{
		{
			files:  []string{"/aaa", "/aaa/bb", "/aaa/cc"},
			check:  "/aaa",
			expect: []string{"/aaa/bb", "/aaa/cc"},
		},
		{
			files:  []string{"/bbb"},
			check:  "/bbb",
			expect: []string{},
		},
	}

	for _, tt := range cases {
		fsmock := FsMockRepository{
			Files: tt.files,
		}
		actual, _ := fsmock.ListFiles(tt.check)
		assert.Equal(t, tt.expect, actual)
	}
}
