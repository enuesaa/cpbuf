package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBrokenSymlink(t *testing.T) {
	fs := FsRepository{}

	cases := []struct {
		path   string
		expect bool
	}{
		{
			path:   "../../testdata/symlink/link.txt",
			expect: false,
		},
		{
			path:   "../../testdata/symlink/source.txt",
			expect: false,
		},
		{
			path:   "../../testdata/symlink-broken/link.txt",
			expect: true,
		},
	}

	for _, tt := range cases {
		actual, err := fs.IsBrokenSymlink(tt.path)
		assert.Equal(t, err, nil)
		assert.Equal(t, tt.expect, actual)
	}
}
