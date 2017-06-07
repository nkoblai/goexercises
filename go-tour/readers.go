package main

import "golang.org/x/tour/reader"

// MyReader represents type for Read method.
type MyReader struct{}

// Read makes infinite stream of the ASCII character 'A' into slice b, returns length of b and nil error.
func (m MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func main() {
	reader.Validate(MyReader{})
}
