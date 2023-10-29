package config

import (
	"testing"

	"github.com/SergeyCherepiuk/togif/pkg/internal"
)

func Test_SingleDashArgs_SpaceAndEqSign(t *testing.T) {
	args := []string{"-c", "-w", "1280", "-h=720", "-i", "input.mp4"}

	actual := parseArgs(args)
	expected := map[string]string{
		"c": "true",
		"w": "1280",
		"h": "720",
		"i": "input.mp4",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_SingleDashArgs_LongNames(t *testing.T) {
	args := []string{"-width", "1280", "-height=720"}

	actual := parseArgs(args)
	expected := map[string]string{
		"width":  "1280",
		"height": "720",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_SingleDashArgs_Decomposition(t *testing.T) {
	args := []string{"-clean"}

	actual := parseArgs(args)
	expected := map[string]string{
		"c": "true",
		"l": "true",
		"e": "true",
		"a": "true",
		"n": "true",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_SingleDashArgs_Override(t *testing.T) {
	args := []string{"-f", "-quality", "80", "-f", "60", "-quality=40"}

	actual := parseArgs(args)
	expected := map[string]string{
		"f":       "60",
		"quality": "40",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_SingleDashArgs_Garbage(t *testing.T) {
	args := []string{"true", "400", "-f", "80", "f", "_value"}

	actual := parseArgs(args)
	expected := map[string]string{
		"f": "80",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_DoubleDashArgs_SpaceAndEqSign(t *testing.T) {
	args := []string{"--frames", "60", "--quality=50", "--compress", "--input", "input.mp4"}

	actual := parseArgs(args)
	expected := map[string]string{
		"frames":   "60",
		"quality":  "50",
		"compress": "true",
		"input":    "input.mp4",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_DoubleDashArgs_Override(t *testing.T) {
	args := []string{"--input", "input.mp4", "--frames", "60", "--input=new_input.avi"}

	actual := parseArgs(args)
	expected := map[string]string{
		"frames": "60",
		"input":  "new_input.avi",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_DoubleDashArgs_Garbage(t *testing.T) {
	args := []string{"somefile.txt", "--input", "input.mp4", "true", "--compress", "--frames", "60", "20"}

	actual := parseArgs(args)
	expected := map[string]string{
		"frames":   "60",
		"compress": "true",
		"input":    "input.mp4",
	}

	internal.ShouldBe[map[string]string](t, actual, expected)
}

func Test_SingleDashArgs_Empty(t *testing.T) {
	args := []string{}

	actual := parseArgs(args)
	expected := make(map[string]string)

	internal.ShouldBe[map[string]string](t, actual, expected)
}
