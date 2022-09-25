package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 2, BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 3, BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))
	})
}
