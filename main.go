package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) <= 1 {
		return ""
	}

	lastCommonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		foundCommon := false
		for j := len(strs[i]); j >= 0; j-- {
			sub := strs[i][:j]

			if strings.Contains(lastCommonPrefix, sub) {
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
