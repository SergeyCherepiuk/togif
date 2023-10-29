package config

import (
	"strings"
)

func parseArgs(args []string) map[string]string {
	options := make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		hasLongPrefix := strings.HasPrefix(arg, "--")
		hasShortPrefix := strings.HasPrefix(arg, "-")
		containsEqSign := strings.Contains(arg, "=")
		nextIsValue := i+1 < len(args) && !strings.HasPrefix(args[i+1], "-")

		if hasLongPrefix && containsEqSign {
			pair := strings.Split(arg, "=")
			options[pair[0][2:]] = pair[1]
			continue
		}

		if hasLongPrefix && nextIsValue {
			options[arg[2:]] = args[i+1]
			i++
			continue
		}

		if hasLongPrefix {
			options[arg[2:]] = "true"
			continue
		}

		if hasShortPrefix && containsEqSign {
			pair := strings.Split(arg, "=")
			options[pair[0][1:]] = pair[1]
			continue
		}

		if hasShortPrefix && nextIsValue {
			options[arg[1:]] = args[i+1]
			i++
			continue
		}

		if hasShortPrefix {
			for _, option := range strings.Split(arg[1:], "") {
				options[option] = "true"
			}
			continue
		}
	}

	return options
}
