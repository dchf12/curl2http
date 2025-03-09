package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func run(r io.Reader, w io.Writer) error {
	input, err := readInput(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return fmt.Errorf("input is empty")
	}

	if !strings.HasPrefix(input, "curl") {
		return fmt.Errorf("input is not curl command")
	}

	request, err := Parse(input)
	if err != nil {
		return fmt.Errorf("failed to parse curl command: %w", err)
	}

	_, err = fmt.Fprint(w, Format(request))
	return err
}

func readInput(r io.Reader) (string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
