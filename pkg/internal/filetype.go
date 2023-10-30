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
	var buf bytes.Buffer
	tee := io.TeeReader(*file, &buf)
	defer func() {
		*file = bytes.NewReader(buf.Bytes())
	}()

	content, err := io.ReadAll(tee)
	if err != nil {
		return "", err
	}

	mimetype := http.DetectContentType(content)
	if !strings.HasPrefix(mimetype, "video") {
		return "", fmt.Errorf("unsupported or invalid video format")
	}

	return strings.TrimPrefix(mimetype, "video/"), nil
}
