package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	lastCommonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		foundCommon := false
		for j := len(lastCommonPrefix); j >= 0; j-- {
			sub := lastCommonPrefix[:j]

			if strings.Contains(strs[i], sub) {
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
 fmt.Printf("Found prefix %s\n", longestCommonPrefix([]string{"flower","flow","flight"}))
}
