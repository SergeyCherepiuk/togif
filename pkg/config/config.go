package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/SergeyCherepiuk/togif/pkg"
	"github.com/SergeyCherepiuk/togif/pkg/internal"
)

var DefaultConfig = Config{
	Input:   bufio.NewReader(os.Stdin),
	Frames:  10,
	Speed:   1.0,
	Verbose: false,
}

type Config struct {
	Input io.Reader

	OutputPath string  `short:"o" long:"output" info:"Path to the output file (destination), if omitted stdout will be used"`
	Frames     uint64  `short:"f" long:"frames" info:"Sets the frames-rate of the resulting GIF image"`
	Speed      float64 `short:"s" long:"speed" info:"Speeds up (s > 1.0) or slows down (s < 1.0) the output GIF"`
	Verbose    bool    `short:"v" long:"verbose" info:"Displays the progress during conversion"`

	Help bool `short:"h" long:"help" info:"Provide information on existing options"`
}

func (c *Config) Validate() error {
	if c.Input == nil {
		return fmt.Errorf(
			"%s: %s: no input stream provided",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE,
		)
	}

	if _, err := os.Stat(c.OutputPath); err == nil {
		return fmt.Errorf(
			"%s: %s: output file already exists: %s",
			pkg.CLI_NAME, pkg.VALIDATION_STAGE, c.OutputPath,
		)
	}

	if _, err := internal.VideoFileType(&c.Input); err != nil {
		return fmt.Errorf("%s: %s: %v", pkg.CLI_NAME, pkg.VALIDATION_STAGE, err)
	}

	return nil
}

func (c *Config) setInputFile(path string) error {
	var file io.Reader
	var err error
	if file, err = os.Open(path); err != nil {
		return err
	}

	if _, err := internal.VideoFileType(&file); err != nil {
		return fmt.Errorf("unsupported or invalid video format")
	}

	c.Input = file
	return nil
}

func (c *Config) setOptions(options map[string]string) error {
	rvs := c.reflectValues()
	for option, value := range options {
		rv, ok := rvs[option]
		if !ok {
			return fmt.Errorf("the -%s option is not found", option)
		}

		if !rv.IsValid() || !rv.CanSet() {
			return fmt.Errorf("cannot set -%s option", option)
		}

		if err := internal.AssertAndSet(&rv, value); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) reflectValues() map[string]reflect.Value {
	rvs := make(map[string]reflect.Value)
	configRt := reflect.TypeOf(*c)
	configRv := reflect.ValueOf(c).Elem()

	for i := 0; i < configRt.NumField(); i++ {
		field := configRt.Field(i)
		rvs[field.Tag.Get("short")] = configRv.Field(i)
		rvs[field.Tag.Get("long")] = configRv.Field(i)
	}

	return rvs
}

func From(args []string) (Config, error) {
	config := DefaultConfig

	if len(args) <= 0 {
		return config, nil
	}

	if err := config.setInputFile(args[len(args)-1]); err == nil {
		args = args[:len(args)-1]
	}

	options := parseArgs(args)
	if err := config.setOptions(options); err != nil {
		return DefaultConfig, fmt.Errorf(
			"%s: %s: %v",
			pkg.CLI_NAME, pkg.CONFIGURATION_STAGE, err,
		)
	}

	if config.OutputPath == "" {
		config.OutputPath = "pipe:1"
	}

	config.Speed = float64(1.0) / config.Speed

	return config, nil
}
