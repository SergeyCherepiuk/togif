package internal

import (
	"fmt"
	"reflect"
	"strconv"
)

// TODO: Unit-test
func AssertAndSet(rv *reflect.Value, value string) error {
	var parsedValue any
	var err error

	// Asserting
	switch rv.Kind() {
	case reflect.String:
		parsedValue = value
	case reflect.Bool:
		parsedValue, err = strconv.ParseBool(value)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		parsedValue, err = strconv.ParseInt(value, 10, 64)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		parsedValue, err = strconv.ParseUint(value, 10, 64)
	case reflect.Float32, reflect.Float64:
		parsedValue, err = strconv.ParseFloat(value, 64)
	default:
		err = fmt.Errorf("parsing of type %s is unsupported", rv.Kind())
	}

	if err != nil {
		return err
	}

	// Setting
	switch pv := parsedValue.(type) {
	case string:
		rv.SetString(pv)
	case bool:
		rv.SetBool(pv)
	case int64:
		rv.SetInt(pv)
	case uint64:
		rv.SetUint(pv)
	case float64:
		rv.SetFloat(pv)
	}

	return nil
}
