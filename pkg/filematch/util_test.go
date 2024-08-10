package filematch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitRmEmpty(t *testing.T) {
	assert.Equal(t, []string{}, splitRmEmpty("", ""))
	assert.Equal(t, []string{}, splitRmEmpty("aaa", "a"))
	assert.Equal(t, []string{"a", "a", "a"}, splitRmEmpty("ababab", "b"))
}
