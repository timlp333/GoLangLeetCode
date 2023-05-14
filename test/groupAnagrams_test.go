package test

import (
	"fmt"
	"leetCode/arraysHashing"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testArray := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	result := arraysHashing.GroupAnagrams(testArray)
	//[["bat"],["nat","tan"],["ate","eat","tea"]]
	// resultArray := [][]string{}
	// resultArray = append(resultArray, []string{"bat"})
	// resultArray = append(resultArray, []string{"nat", "tan"})
	// resultArray = append(resultArray, []string{"ate", "eat", "tea"})
	fmt.Print(result[0][0])
	if result[0][0] != "bat" {
		t.Errorf("Expected result[0][0] to be bat, but got %s instead", result[0][0])
	}
}
func TestGroupAnagrams2(t *testing.T) {
	testArray := []string{""}

	result := arraysHashing.GroupAnagrams(testArray)
	//[["bat"],["nat","tan"],["ate","eat","tea"]]
	// resultArray := [][]string{}
	// resultArray = append(resultArray, []string{"bat"})
	// resultArray = append(resultArray, []string{"nat", "tan"})
	// resultArray = append(resultArray, []string{"ate", "eat", "tea"})
	fmt.Print(result[0][0])
	if result[0][0] != "" {
		t.Errorf("Expected result[0][0] to be \"\", but got %s instead", result[0][0])
	}
}
