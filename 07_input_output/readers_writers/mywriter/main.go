package main

import (
	"fmt"
	"io"
)

type NWriter struct {
	data string
	size int
}

func (w *NWriter) Write(p []byte) (int, error) {
	if p == nil {
		return -1, fmt.Errorf("target array is nil")
	}
	if len(p) == 0 {
		return 0, io.EOF
	}

	n := w.size
	var err error = nil
	if len(p) <= n {
		n = len(p)
	} else {
		err = fmt.Errorf("p larger than size")
	}

	w.data = w.data + string(p[0:n])
	return n, err
}

func main() {
	msg := []byte("the world with Go!!!")

	mw := NWriter{"Save ", 6}
	i := 0
	var err error
	for err == nil && i < len(msg) {
		n, err := mw.Write(msg[i:])
		fmt.Printf("Written %d error %v --> %s\n", n, err, mw.data)
		i = i + n
	}
}
