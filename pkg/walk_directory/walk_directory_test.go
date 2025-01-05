package walk_directory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkDirectory(t *testing.T) {
	// Test WalkDirectory
	totalSize, totalFiles, err := WalkDirectory("test_data")

	assert.Nil(t, err)
	assert.Equal(t, uint64(78), totalSize)
	assert.Equal(t, uint64(4), totalFiles)
}
