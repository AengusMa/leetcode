package leetcode049

import (
	"fmt"
	"testing"
)

func Test049(t *testing.T) {
	// [
	// 	["ate","eat","tea"],
	// 	["nat","tan"],
	// 	["bat"]
	// ]
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"dad", "day", "dad", "dat", "dad", "day"}))
}
