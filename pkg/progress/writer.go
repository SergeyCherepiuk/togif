package progress

import "io"

type WriteCloser struct {
	io.WriteCloser
	ch chan<- int
}

func NewWriteCloser(wc io.WriteCloser) (io.WriteCloser, <-chan int) {
	ch := make(chan int)
	return WriteCloser{WriteCloser: wc, ch: ch}, ch
}

func (wwp WriteCloser) Write(b []byte) (int, error) {
	n, err := wwp.WriteCloser.Write(b)
	wwp.ch <- n
	return n, err
}

func (wwp WriteCloser) Close() error {
	close(wwp.ch)
	return wwp.WriteCloser.Close()
}
