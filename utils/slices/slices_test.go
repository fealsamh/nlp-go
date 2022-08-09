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

func TestEqual(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.Equal(false, Equal([]int{1, 2, 3}, []int{1, 2}))
	assert.Equal(false, Equal([]int{1, 2, 3}, []int{1, 2, 4}))
}
