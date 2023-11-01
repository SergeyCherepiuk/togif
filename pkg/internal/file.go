package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Unit-test
func VideoFileType(file *io.Reader) (string, error) {
	b, err := getBytes(file)
	if err != nil {
		return "", err
	}

	mimetype := http.DetectContentType(b)
	if !strings.HasPrefix(mimetype, "video") {
		return "", fmt.Errorf("unsupported or invalid video format")
	}

	return strings.TrimPrefix(mimetype, "video/"), nil
}

// NOTE: Inefficient workaround
// TODO: Unit-test
func FileSize(file *io.Reader) (int, error) {
	b, err := getBytes(file)
	return len(b), err
}

func getBytes(r *io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(*r, &buf)

	b, err := io.ReadAll(tee)
	*r = bytes.NewReader(buf.Bytes())

	return b, err
}
