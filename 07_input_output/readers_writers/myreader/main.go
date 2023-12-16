package main

import (
	"fmt"
	"io"
)

type MReader struct {
	data string
	from int
}

func NewMReader(data string) *MReader {
	return &MReader{
		data: data,
		from: 0,
	}
}

func (r *MReader) Read(p []byte) (int, error) {
	if p == nil {
		return -1, fmt.Errorf("target array is nil")
	}

	if len(r.data) <= 0 || r.from == len(r.data) {
		return 0, io.EOF
	}

	n := len(r.data) - r.from
	if len(p) < n {
		n = len(p)
	}

	for i := 0; i < n; i++ {
		b := byte(r.data[r.from])
		p[i] = b
		r.from++
	}

	if len(r.data) == r.from {
		return n, io.EOF
	}

	return n, nil
}

func main() {
	target := make([]byte, 5)

	empty := NewMReader("")
	n, err := empty.Read(target)
	fmt.Printf("Read %d: Error: %v\n", n, err)
	r := NewMReader("Save the world with Go!!!")
	n, err = r.Read(target)
	for err == nil {
		fmt.Printf("Read %d: Error: %v -> %s\n", n, err, target)
		n, err = r.Read(target)
	}
	fmt.Printf("Read %d: Error: %v -> %s\n", n, err, target)
}
