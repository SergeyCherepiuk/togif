package internal

import (
	"strings"
)

// TODO: Optimize and unit test
func Tabulate(lines [][]string, extraSpaces int) []string {
	if len(lines) <= 0 {
		return make([]string, 0)
	}

	var maxPartsCount int
	for _, line := range lines {
		if len(line) > maxPartsCount {
			maxPartsCount = len(line)
		}
	}

	results := make([]string, len(lines))

	for i := 0; i < maxPartsCount; i++ {
		var maxPartLen int
		for _, line := range lines {
			if i < len(line) && len(line[i]) > maxPartLen {
				maxPartLen = len(line[i])
			}
		}

		for j, line := range lines {
			if i < len(line) {
				spaces := strings.Repeat(" ", maxPartLen-len(line[i])+extraSpaces)
				results[j] = results[j] + line[i] + spaces
			}
		}
	}

	return results
}
