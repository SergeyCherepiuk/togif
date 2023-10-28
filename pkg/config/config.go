package config

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/SergeyCherepiuk/togif/pkg"
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

func from(flags map[string]string) (Config, error) {
	config := DefaultConfig
	rvs := make(map[string]reflect.Value)

	configRt := reflect.TypeOf(config)
	configRv := reflect.ValueOf(&config).Elem()

	for i := 0; i < configRt.NumField(); i++ {
		field := configRt.Field(i)
		rvs[field.Tag.Get("short")] = configRv.Field(i)
		rvs[field.Tag.Get("long")] = configRv.Field(i)
	}

	for flag, value := range flags {
		rv, ok := rvs[flag]
		if !ok {
			return DefaultConfig, fmt.Errorf(
				"%s: %s: failed to set -%s flag",
				pkg.CLI_NAME, pkg.CONF_STAGE, flag,
			)
		}

		if !rv.IsValid() || !rv.CanSet() {
			return DefaultConfig, fmt.Errorf(
				"%s: %s: cannot set -%s flag",
				pkg.CLI_NAME, pkg.CONF_STAGE, flag,
			)
		}

		parseAndSet(&rv, value)
	}

	return config, nil
}

func parseAndSet(rv *reflect.Value, value string) {
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
		panic(fmt.Sprintf("%s: %s: %v", pkg.CLI_NAME, pkg.CONF_STAGE, err))
	}
}
