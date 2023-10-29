package internal

import (
	"strings"
)

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

	result := make([]string, len(lines))

	for i := 0; i < maxPartsCount; i++ {
		var maxPartLength int
		for _, line := range lines {
			if i < len(line) && len(line[i]) > maxPartLength {
				maxPartLength = len(line[i])
			}
		}

		for j := range lines {
			if i < len(lines[j]) {
				result[j] = result[j] + lines[j][i]
			}

			if i+1 < len(lines[j]) {
				spaces := strings.Repeat(" ", maxPartLength-len(lines[j][i])+extraSpaces)
				result[j] = result[j] + spaces
			}
		}
	}

	return result
}
