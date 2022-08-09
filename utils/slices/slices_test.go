package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Contains([]int{1, 2, 3}, 2))
	assert.Equal(false, Contains([]int{1, 2, 3}, 4))
}
