package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo2(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 0, Fibo2(0))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 55, Fibo2(10))
	})
}
