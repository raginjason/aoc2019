package util

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

const baseDir = "/data/day"

// FileInput grab input data for a given day and return a scanner
func FileInput(day int) ([]string, error) {
	var lines []string
	filename := filepath.Join(baseDir, strconv.Itoa(day), "input")

	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, err
}

// FileInputReader grab input data for a given day and return a reader
func FileInputReader(day int) (io.Reader, error) {
	filename := filepath.Join(baseDir, strconv.Itoa(day), "input")

	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	return fh, err
}
