package internal

import (
	"reflect"
	"testing"
)

func ShouldBe[T any](t *testing.T, actual, expected T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed\nactual: %+v\nexpected: %+v\n", actual, expected)
	}
}
