package progress

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Display(ch <-chan int, size int) {
	fmt.Fprintln(os.Stderr)

	var bytesRead int

	var start time.Time
	var eta time.Duration
	go func() {
		for {
			ratio := float64(size-bytesRead) / float64(bytesRead)
			eta = time.Duration(ratio*time.Since(start).Seconds()) * time.Second
			time.Sleep(time.Second)
		}
	}()

	for n := range ch {
		if start.IsZero() {
			start = time.Now()
		}

		bytesRead += n
		write(float32(bytesRead)/float32(size)*100, eta)
	}
	write(100.0, 0)
}

func write(progress float32, eta time.Duration) {
	fmt.Fprintf(os.Stderr, "\033[1A\033[K%s\n", format(progress, eta, 50))
}

func format(progress float32, eta time.Duration, length int) string {
	f := float32(100.0) / float32(length)
	done := strings.Repeat("=", int(progress/f))
	left := strings.Repeat(" ", length-int(progress/f))
	return fmt.Sprintf("Progress: [%s%s] %6.2f%% | ETA: %v", done, left, progress, eta)
}
