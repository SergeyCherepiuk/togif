package internal

import (
	"reflect"
	"testing"
)

func Test_AssertAndSet_String(t *testing.T) {
	value := "old"

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "new")

	ShouldBe[string](t, value, "new")
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Bool_True(t *testing.T) {
	value := false

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "true")

	ShouldBe[bool](t, value, true)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Bool_False(t *testing.T) {
	value := true

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "false")

	ShouldBe[bool](t, value, false)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Bool_String(t *testing.T) {
	value := true

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "text")

	ShouldBe[bool](t, value, true)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Bool_Int(t *testing.T) {
	value := true

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "123")

	ShouldBe[bool](t, value, true)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Int64_Positive(t *testing.T) {
	value := int64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "100")

	ShouldBe[int64](t, value, 100)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Int64_Negative(t *testing.T) {
	value := int64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "-100")

	ShouldBe[int64](t, value, -100)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Int64_Bool(t *testing.T) {
	value := int64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "true")

	ShouldBe[int64](t, value, 50)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Uint64_Positive(t *testing.T) {
	value := uint64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "100")

	ShouldBe[uint64](t, value, 100)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Uint64_Negative(t *testing.T) {
	value := uint64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "-100")

	ShouldBe[uint64](t, value, 50)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Uint64_Float(t *testing.T) {
	value := uint64(50)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "100.0")

	ShouldBe[uint64](t, value, 50)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Float64_Whole(t *testing.T) {
	value := float64(1.0)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "2")

	ShouldBe[float64](t, value, 2.0)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Float64_Decimal(t *testing.T) {
	value := float64(1.0)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "2.999")

	ShouldBe[float64](t, value, 2.999)
	ShouldBe[error](t, err, nil)
}

func Test_AssertAndSet_Float64_Bool(t *testing.T) {
	value := float64(1.0)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "false")

	ShouldBe[float64](t, value, 1.0)
	ShouldNotBe[error](t, err, nil)
}

func Test_AssertAndSet_Float64_String(t *testing.T) {
	value := float64(1.0)

	rv := reflect.ValueOf(&value).Elem()
	err := AssertAndSet(&rv, "text")

	ShouldBe[float64](t, value, 1.0)
	ShouldNotBe[error](t, err, nil)
}
