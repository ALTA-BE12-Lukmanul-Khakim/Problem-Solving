package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoneyChange(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []int{100, 20, 1, 1, 1}, MoneyChange(123))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, []int{10000, 5000, 200, 100, 20, 1}, MoneyChange(15321))
	})
}
