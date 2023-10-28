package config

import (
	"strings"
)

// TODO: Unit-test

func Parse(args []string) Config {
	config := DefaultConfig
	return config
}

func parseArgs(args []string) map[string]string {
	pairs := make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		hasLongPrefix := strings.HasPrefix(arg, "--")
		hasShortPrefix := strings.HasPrefix(arg, "-")
		containsEqSign := strings.Contains(arg, "=")
		nextIsValue := i+1 < len(args) && !strings.HasPrefix(args[i+1], "-")

		if hasLongPrefix && containsEqSign {
			pair := strings.Split(arg, "=")
			pairs[pair[0][2:]] = pair[1]
			continue
		}

		if hasLongPrefix && nextIsValue {
			pairs[arg[2:]] = args[i+1]
			i++
			continue
		}

		if hasLongPrefix {
			pairs[arg[2:]] = "true"
			continue
		}

		if hasShortPrefix && containsEqSign {
			pair := strings.Split(arg, "=")
			pairs[pair[0][1:]] = pair[1]
			continue
		}

		if hasShortPrefix && nextIsValue {
			pairs[arg[1:]] = args[i+1]
			i++
			continue
		}

		if hasShortPrefix {
			for _, flag := range strings.Split(arg[1:], "") {
				pairs[flag] = "true"
			}
			continue
		}
	}

	return pairs
}
