package handler

import (
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/togif/pkg"
	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/help"
	"github.com/SergeyCherepiuk/togif/pkg/validation"
)

func Handle(args []string) {
	config, err := config.From(args)
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

	if config.Input == nil {
		return fmt.Errorf(
			"%s: %s: no input stream provided",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE,
		)
	}

	if config.Output == nil {
		return fmt.Errorf(
			"%s: %s: no output stream provided",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE,
		)
	}

	if is, err := validation.IsVideoFile(&config.Input); !is || err != nil {
		return fmt.Errorf(
			"%s: %s: unsupported or invalid video format",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE,
		)
	}

	return nil
}
