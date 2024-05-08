package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	// Test case 1: Successful test
	result := Sum(2, 3)
	assert.Equal(t, 5, result, "The sum of 2 and 3 should be 5")

	// Test case 2: Failed test
	// result = Sum(2, 3)
	// assert.Equal(t, 6, result, "The sum of 2 and 3 should be 6")
}
