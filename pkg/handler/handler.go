package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/SergeyCherepiuk/togif/pkg"
	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/help"
)

func Handle(args []string) {
	config, err := config.From(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		help.Display(os.Stderr)
		return
	}

	if config.Help {
		help.Display(os.Stdout)
		return
	}

	if err := validate(config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		help.Display(os.Stderr)
		return
	}

	fmt.Printf("Execution continues: %+v\n", config)
}

func validate(config config.Config) error {
	if config.Quality > 100 {
		return fmt.Errorf(
			"%s: %s: invalid quality value: should be in [0, 100] range",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE,
		)
	}

	if config.Input != nil {
		if buf, err := io.ReadAll(config.Input); err == nil {
			if !strings.HasPrefix(http.DetectContentType(buf), "video") {
				return fmt.Errorf(
					"%s: %s: unsupported or invalid video format",
					pkg.CLI_NAME, pkg.VALIDATION_STAGE,
				)
			}
		}
	}

	return nil
}
