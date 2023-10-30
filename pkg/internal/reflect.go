package internal

import (
	"reflect"
	"strconv"
)

// TODO: Unit-test
func AssertAndSet(rv *reflect.Value, value string) error {
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

	return err
}
