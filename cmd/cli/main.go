package main

import (
	"os"

	"github.com/SergeyCherepiuk/togif/pkg/handler"
)

func main() {
	handler.Handle(os.Args[1:])
}
