package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	fileName, err := GetFileName()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = OpenFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GetFileName() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New("No file name was provided")
	}

	return os.Args[1], nil
}

func OpenFile(fileName string) error {
	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	io.Copy(os.Stdout, file)
	return nil
}
