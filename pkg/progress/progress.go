package progress

import (
	"fmt"
	"os"
	"strings"
)

func Display(ch <-chan int, size int) {
	fmt.Fprintln(os.Stderr)

	var bytesRead int
	for n := range ch {
		bytesRead += n
		write(float32(bytesRead) / float32(size) * 100)
	}
	write(100.0)
}

func write(percent float32) {
	fmt.Fprintf(os.Stderr, "\033[1A\033[K%s\n", format(percent, 50))
}

func format(percent float32, length int) string {
	f := float32(100) / float32(length)
	done := strings.Repeat("=", int(percent/f))
	left := strings.Repeat(" ", length-int(percent/f))
	return fmt.Sprintf("Progress: [%s%s] %6.2f%%", done, left, percent)
}
