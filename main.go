package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) <= 1 {
		return ""
	}

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	lastCommonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		foundCommon := false
		nLen := int(math.Min(float64(len(lastCommonPrefix)), float64(len(strs[i]))))
		for j := nLen; j >= 0; j-- {
			sub := strs[i][:j]

			if lastCommonPrefix[:len(sub)] == sub {
				lastCommonPrefix = sub
				foundCommon = true
				break
			}
		}

		if !foundCommon {
			return ""
		}
	}

	return lastCommonPrefix
}

func main() {
 fmt.Printf("Found prefix %s\n", longestCommonPrefix([]string{"c","acc","ccc"}))
}
