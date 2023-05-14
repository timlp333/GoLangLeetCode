package test

import (
	"leetCode/arraysHashing"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	result := arraysHashing.IsAnagram("anagram", "nagaram")
	if !result {
		t.Errorf("Expected result to be false, but got %t instead", result)
	}
}
