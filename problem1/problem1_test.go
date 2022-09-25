package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleEquation(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "no solution", SimpleEquation(1, 2, 3))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, "1 2 3", SimpleEquation(6, 6, 14))
	})
}
