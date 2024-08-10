package filematch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	assert.Equal(t, true, Is("a.txt", "a.txt"))
	assert.Equal(t, true, Is("a.txt", "."))
	assert.Equal(t, true, Is("a.txt", "*"))
	assert.Equal(t, true, Is("a.txt", "a*"))
	assert.Equal(t, true, Is("a.txt", "*.txt"))
	assert.Equal(t, true, Is("a/bb.txt", "a"))
	assert.Equal(t, true, Is("a/bb/cc.txt", "a"))
	assert.Equal(t, true, Is("a/bb/cc.txt", "a*"))
	assert.Equal(t, true, Is("ab/bb/cc.txt", "a*"))
	assert.Equal(t, true, Is("ab/bb/cc.txt", "ab"))
	assert.Equal(t, true, Is("ab/bb/cc.txt", "ab/bb"))
	assert.Equal(t, true, Is("ab/bb/cc.txt", "ab/bb/cc.txt"))

	assert.Equal(t, false, Is("a.txt", "b.txt"))
	assert.Equal(t, false, Is("a.txt", ""))
	assert.Equal(t, false, Is("a.txt", "b*"))
	assert.Equal(t, false, Is("a.txt", "a.txta"))
	assert.Equal(t, false, Is("ab/bb/cc.txt", "a"))
	assert.Equal(t, false, Is("ab/bb/cc.txt", "*a"))
}

func TestIsTextMatch(t *testing.T) {
	assert.Equal(t, true, isTextMatch("a.txt", "a.txt"))
	assert.Equal(t, true, isTextMatch("a.txt", "*.txt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a.*xt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a*xt"))
	assert.Equal(t, true, isTextMatch("a.txt", "a*"))

	assert.Equal(t, false, isTextMatch("ab", "a"))
	assert.Equal(t, false, isTextMatch("ab", "*a"))
	assert.Equal(t, false, isTextMatch("a.txt", "b*"))
	assert.Equal(t, false, isTextMatch("a.txt", "*c"))
}
