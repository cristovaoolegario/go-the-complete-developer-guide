package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("Should return an error message When there's no args to program", func(t *testing.T) {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()
		os.Args = []string{"cmd"}

		_, err := GetFileName()

		if err.Error() != "No file name was provided" {
			t.Errorf("Expected to thronw an error!")
		}
	})

	t.Run("Should return file name When args has been passed to program", func(t *testing.T) {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()
		os.Args = []string{"cmd", "TestFileName"}

		fileName, err := GetFileName()

		if fileName != "TestFileName" && err != nil {
			t.Errorf("Mismatch file name!")
		}
	})

	t.Run("Should return an error message When there's no file to open", func(t *testing.T) {
		err := OpenFile("NotExistingFile.txt")

		if err.Error() != "open NotExistingFile.txt: no such file or directory" {
			t.Error("Expected to thronw an error! Got:", err)
		}
	})
}
