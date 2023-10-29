package internal

import (
	"testing"
)

func Test_Tabulate_EmptyMatrix_OneSpace(t *testing.T) {
	lines := [][]string{}

	actual := Tabulate(lines, 1)
	expected := []string{}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_EmptyMatrix_FiveSpaces(t *testing.T) {
	lines := [][]string{}

	actual := Tabulate(lines, 5)
	expected := []string{}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_EmptyLines_OneSpace(t *testing.T) {
	lines := [][]string{
		{""},
		{""},
	}

	actual := Tabulate(lines, 1)
	expected := []string{"", ""}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_EmptyLines_FiveSpaces(t *testing.T) {
	lines := [][]string{
		{""},
		{""},
	}

	actual := Tabulate(lines, 5)
	expected := []string{"", ""}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_RectangularMatrix_2x2_OneSpace(t *testing.T) {
	lines := [][]string{
		{"first", "second value"},
		{"third value", "forth"},
	}

	actual := Tabulate(lines, 1)
	expected := []string{
		"first       second value",
		"third value forth",
	}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_RectangularMatrix_2x2_FiveSpaces(t *testing.T) {
	lines := [][]string{
		{"first", "second value"},
		{"third value", "forth"},
	}

	actual := Tabulate(lines, 5)
	expected := []string{
		"first           second value",
		"third value     forth",
	}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_RectangularMatrix_2x3_OneSpace(t *testing.T) {
	lines := [][]string{
		{"first value", "second value", "third"},
		{"forth value", "fifth", "sixth value"},
	}

	actual := Tabulate(lines, 1)
	expected := []string{
		"first value second value third",
		"forth value fifth        sixth value",
	}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_RectangularMatrix_2x3_FiveSpaces(t *testing.T) {
	lines := [][]string{
		{"first value", "second value", "third"},
		{"forth value", "fifth", "sixth value"},
	}

	actual := Tabulate(lines, 5)
	expected := []string{
		"first value     second value     third",
		"forth value     fifth            sixth value",
	}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_JaggedMatrix_OneSpaces(t *testing.T) {
	lines := [][]string{
		{"first value", "second value", "third"},
		{"forth"},
		{"fifth value", "sixth value", "seventh"},
	}

	actual := Tabulate(lines, 1)
	expected := []string{
		"first value second value third",
		"forth",
		"fifth value sixth value  seventh",
	}

	ShouldBe[[]string](t, actual, expected)
}

func Test_Tabulate_JaggedMatrix_FiveSpaces(t *testing.T) {
	lines := [][]string{
		{"first value", "second value", "third"},
		{"forth"},
		{"fifth value", "sixth value", "seventh"},
	}

	actual := Tabulate(lines, 5)
	expected := []string{
		"first value     second value     third",
		"forth",
		"fifth value     sixth value      seventh",
	}

	ShouldBe[[]string](t, actual, expected)
}
