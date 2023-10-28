package config

import (
	"fmt"
	"reflect"
	"strconv"
)

var DefaultConfig = Config{
	Frames:  10,
	Quality: 80,
}

type Config struct {
	Frames    uint64 `short:"f" long:"frames"`  // frames per second
	Quality   uint8  `short:"q" long:"quality"` // quality, [0%, 100%]
	StartMs   uint64 `short:"s" long:"start"`   // start point (in milliseconds)
	EndMs     uint64 `short:"e" long:"end"`     // end point (in milliseconds)
	IFilename string `short:"i" long:"input"`   // input file/filepath (source)
	OFilename string `short:"o" long:"output"`  // output file/filepath (destination)
}

func From(args map[string]string) (Config, error) {
	// TODOs:
	//  1. Create a map (map[string]reflect.Value) for short and long tags
	//  2. Iterate through args and get a proper reflect.Value
	return DefaultConfig, nil
}

func parseAndSet(rv reflect.Value, value string) {
	// TODO: Check if value can be set

	var err error
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(value)

	case reflect.Bool:
		var b bool
		b, err = strconv.ParseBool(value)
		rv.SetBool(b)

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		var i int64
		i, err = strconv.ParseInt(value, 10, 64)
		rv.SetInt(i)
	}

	if err != nil {
		panic(fmt.Sprintf("togif: parsing arguments: %v", err))
	}
}
