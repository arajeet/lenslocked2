package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder
	WriteGreeting("Jon", &sb)
}

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func WriteGreeting(name string, w Writer) error {
	greeting := Greeting(name)
	// We can access Write here because it is defined by the Writer interface,
	// but we CANNOT access the Builder.WriteString method here even if we passed
	// in a strings.Builder because WriteString is not defined as part of the
	// Writer interface.
	_, err := w.Write([]byte(greeting))
	if err != nil {
		return err
	}
	return nil
}

type Writer interface {
	Write([]byte) (int, error)
}
