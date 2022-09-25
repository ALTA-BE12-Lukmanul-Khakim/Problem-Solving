package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibo(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 0, Fibo(0))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 55, Fibo(10))
	})
}
