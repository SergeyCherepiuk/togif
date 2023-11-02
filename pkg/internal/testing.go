package internal

import (
	"reflect"
	"testing"
)

func ShouldBe[T any](t *testing.T, actual, expected T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual should be equal to expected\nactual: %+v\nexpected: %+v\n", actual, expected)
	}
}

func ShouldNotBe[T any](t *testing.T, actual, expected T) {
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual should not be equal to expected\nactual: %+v\nexpected: %+v\n", actual, expected)
	}
}
