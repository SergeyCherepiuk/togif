package handler

import (
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/gif"
	"github.com/SergeyCherepiuk/togif/pkg/help"
)

func Handle(args []string) {
	var err error
	defer func() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			help.Display(os.Stderr)
		}
	}()

	config, err := config.From(args)
	if err != nil {
		return
	}

	if config.Help {
		help.Display(os.Stdout)
		return
	}

	if err = config.Validate(); err != nil {
		return
	}

	err = gif.Convert(config)
}
