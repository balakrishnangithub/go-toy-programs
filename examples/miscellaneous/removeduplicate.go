package main

import (
	"bufio"
	"errors"
	"os"
)

// rmDupeGenFile reads and generates a file without duplicate lines.
func rmDupeGenFile(srcPath string, destPath string) error {
	if srcPath == destPath {
		return errors.New("srcPath is equal to destPath which is not accepted")
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer destFile.Close()

	dupeCount := make(map[string]int)

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriterSize(destFile, 2048) // defaultBufSize is 4096 bytes

	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := dupeCount[line]; ok {
			dupeCount[line]++
		} else {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
			dupeCount[line] = 0
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}
	return err
}
