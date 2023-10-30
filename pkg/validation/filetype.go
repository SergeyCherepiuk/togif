package validation

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

// TODO: Unit-test
func IsVideoFile(file *io.Reader) (bool, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(*file, &buf)
	defer func() {
		*file = bytes.NewReader(buf.Bytes())
	}()

	content, err := io.ReadAll(tee)
	if err != nil {
		return false, err
	}

	return strings.HasPrefix(http.DetectContentType(content), "video"), nil
}
