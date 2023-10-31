package handler

import (
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/gif"
	"github.com/SergeyCherepiuk/togif/pkg/help"
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

	if err := config.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		help.Display(os.Stderr)
		return
	}

	if err := gif.Convert(config); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
