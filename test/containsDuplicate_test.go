package test

import (
	"leetCode/arraysHashing"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	result := arraysHashing.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if !result {
		t.Errorf("Expected result to be true, but got %t instead", result)
	}
}
