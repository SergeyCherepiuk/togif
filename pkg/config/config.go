package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/SergeyCherepiuk/togif/pkg"
)

var DefaultConfig = Config{
	Input:   bufio.NewReader(os.Stdin),
	Frames:  10,
	Quality: 80,
}

type Config struct {
	Input io.Reader

	Output  string `short:"o" long:"output" info:"Path to the output file (destination), if omitted stdout will be used"`
	Frames  uint64 `short:"f" long:"frames" info:"Sets the frames-rate of the resulting GIF image"`
	Quality uint8  `short:"q" long:"quality" info:"Sets the quality of the resulting GIF image, should be an integer number in [0, 100] range"`

	Help bool `short:"h" long:"help" info:"Provide information on existing options"`
}

func From(args []string) (Config, error) {
	config := DefaultConfig

	if len(args) <= 0 {
		return config, nil
	}

	// Setting an input file
	inputPath := args[len(args)-1]
	if info, err := os.Stat(inputPath); err == nil && !info.IsDir() {
		if inputFile, err := os.Open(inputPath); err == nil {
			config.Input = inputFile
		}
		args = args[:len(args)-1]
	}

	// Parsing and setting the options
	options := parseArgs(args)

	rvs := make(map[string]reflect.Value)
	configRt := reflect.TypeOf(config)
	configRv := reflect.ValueOf(&config).Elem()

	for i := 0; i < configRt.NumField(); i++ {
		field := configRt.Field(i)
		rvs[field.Tag.Get("short")] = configRv.Field(i)
		rvs[field.Tag.Get("long")] = configRv.Field(i)
	}

	for option, value := range options {
		rv, ok := rvs[option]
		if !ok {
			return DefaultConfig, fmt.Errorf(
				"%s: %s: failed to set -%s option",
				pkg.CLI_NAME, pkg.CONFIGURATION_STAGE, option,
			)
		}

		if !rv.IsValid() || !rv.CanSet() {
			return DefaultConfig, fmt.Errorf(
				"%s: %s: cannot set -%s option",
				pkg.CLI_NAME, pkg.CONFIGURATION_STAGE, option,
			)
		}

		if err := assertAndSet(&rv, value); err != nil {
			return DefaultConfig, err
		}
	}

	return config, nil
}

func assertAndSet(rv *reflect.Value, value string) error {
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

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		var u uint64
		u, err = strconv.ParseUint(value, 10, 64)
		rv.SetUint(u)
	}

	if err != nil {
		return fmt.Errorf("%s: %s: %v", pkg.CLI_NAME, pkg.CONFIGURATION_STAGE, err)
	}

	return nil
}
