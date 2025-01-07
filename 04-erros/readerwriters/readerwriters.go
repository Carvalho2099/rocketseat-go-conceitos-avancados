package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Ahoy"
	reader := strings.NewReader(str)

	buffer := make([]byte, 2)
	n, err := reader.Read(buffer)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(n)
	fmt.Println(string(buffer[:n]))

	writer := MyWriter{}
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		_, _ = writer.Writer(buffer[:n])
	}
}

type MyWriter struct{}

func (MyWriter) Writer(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
