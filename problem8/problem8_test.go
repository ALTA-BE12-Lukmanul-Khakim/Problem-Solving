package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "VI", RomanNumerals(6))
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, "XXIII", RomanNumerals(23))
	})
}
