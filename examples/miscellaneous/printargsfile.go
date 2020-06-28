package main

import (
	"errors"
	"os"

	myutils "github.com/bkmagnetron/go-toy-programs/pkg/utils"
)

func printArgsFile() (err error) {
	if len(os.Args) != 2 {
		return errors.New("args error")
	}
	srcPath := os.Args[1]

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	myutils.CoutAll(srcFile)

	return err
}
