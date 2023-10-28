package main

import (
	"fmt"
	"os"

	"github.com/SergeyCherepiuk/togif/pkg/config"
)

func main() {
	config := config.MustParse(os.Args[1:])
	fmt.Printf("%+v\n", config)
}
