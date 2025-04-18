package stdin

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var (
	ErrEmptyInput     = errors.New("empty input provided")
	ErrUnhandledInput = errors.New("unhandled input mode")
)

func Read() (string, error) {
	return ReadFrom(os.Stdin)
}

func ReadFrom(in *os.File) (string, error) {
	stat, err := in.Stat()
	if err != nil {
		return "", err
	}
	// input is piped
	if stat.Mode()&os.ModeNamedPipe != 0 {
		return untilEOF(in)
	}
	// input is a file
	if stat.Mode().IsRegular() {
		return untilEOF(in)
	}
	// input is interactive
	if stat.Mode()&os.ModeDevice != 0 {
		return untilNewline(in)
	}
	return "", ErrUnhandledInput
}

func untilEOF(f *os.File) (string, error) {
	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(string(bytes))
	return ensureNotEmpty(value)
}

func untilNewline(f *os.File) (string, error) {
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	value := strings.TrimSpace(scanner.Text())
	return ensureNotEmpty(value)
}

func ensureNotEmpty(value string) (string, error) {
	if value == "" {
		return "", ErrEmptyInput
	}
	return value, nil
}
