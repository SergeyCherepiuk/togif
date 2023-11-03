package internal

import (
	"bytes"
	"io"
	"testing"
)

func Test_VideoFileType_MP4(t *testing.T) {
	var file io.Reader = bytes.NewReader(
		[]byte{
			0x00, 0x00, 0x00, 0x0C,
			0x66, 0x74, 0x79, 0x70,
			0x6d, 0x70, 0x34, 0xff,
		},
	)

	actual, err := VideoFileType(&file)
	expected := "mp4"

	ShouldBe[string](t, actual, expected)
	ShouldBe[error](t, err, nil)
}

func Test_VideoFileType_WebM(t *testing.T) {
	var file io.Reader = bytes.NewReader(
		[]byte{0x1a, 0x45, 0xdf, 0xa3},
	)

	actual, err := VideoFileType(&file)
	expected := "webm"

	ShouldBe[string](t, actual, expected)
	ShouldBe[error](t, err, nil)
}

func Test_VideoFileType_AVI(t *testing.T) {
	var file io.Reader = bytes.NewReader(
		[]byte{
			0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x40,
			0x41, 0x56, 0x49, 0x20, 0x4C, 0x49, 0x53, 0x54,
		},
	)

	actual, err := VideoFileType(&file)
	expected := "avi"

	ShouldBe[string](t, actual, expected)
	ShouldBe[error](t, err, nil)
}
