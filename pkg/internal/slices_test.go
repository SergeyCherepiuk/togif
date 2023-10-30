package internal

import (
	"testing"
)

func Test_Filter_Odd_Empty(t *testing.T) {
	slice := []int{}

	actual := Filter[int](slice, func(i int) bool { return i%2 == 1 })
	expected := []int{}

	ShouldBe[[]int](t, actual, expected)
}

func Test_Filter_Odd_0_9(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	actual := Filter[int](slice, func(i int) bool { return i%2 == 1 })
	expected := []int{1, 3, 5, 7, 9}

	ShouldBe[[]int](t, actual, expected)
}

func Test_Filter_Uppercase_Empty(t *testing.T) {
	slice := []byte{}

	actual := Filter[byte](slice, func(b byte) bool { return 'A' <= b && b <= 'Z' })
	expected := []byte{}

	ShouldBe[[]byte](t, actual, expected)
}

func Test_Filter_Uppercase_Hello(t *testing.T) {
	slice := []byte("HeLlO")

	actual := Filter[byte](slice, func(b byte) bool { return 'A' <= b && b <= 'Z' })
	expected := []byte{'H', 'L', 'O'}

	ShouldBe[[]byte](t, actual, expected)
}

type window struct {
	width  int
	height int
}

func Test_Filter_WindowWideEnough_Empty(t *testing.T) {
	slice := []window{}

	actual := Filter[window](slice, func(w window) bool { return w.width > 2000 })
	expected := []window{}

	ShouldBe[[]window](t, actual, expected)
}

func Test_Filter_WindowWideEnough_NotEmpty(t *testing.T) {
	slice := []window{
		{1280, 720}, {8200, 1820}, {180, 7200}, {2040, 420}, {0, 2510},
	}

	actual := Filter[window](slice, func(w window) bool { return w.width > 2000 })
	expected := []window{{8200, 1820}, {2040, 420}}

	ShouldBe[[]window](t, actual, expected)
}
