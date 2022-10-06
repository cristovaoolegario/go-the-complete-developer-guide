package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetTriangleArea(t *testing.T) {
	tc := []struct {
		base   float64
		height float64
		area   float64
	}{{1, 1, 0.5}, {1, 2, 1}, {2, 1, 1}, {3, 14, 21}, {5, 4, 10}}

	for _, test := range tc {
		tr := triangle{base: test.base, height: test.height}
		area := tr.getArea()
		if area != test.area {
			t.Errorf("Expected triagle area %v to be %v", area, test.area)
		}
	}
}

func TestGetSquareArea(t *testing.T) {
	tc := []struct {
		sideLength float64
		area       float64
	}{{1, 1}, {2, 4}, {3, 9}, {4, 16}, {5, 25}}

	for _, test := range tc {
		sq := square{sideLength: test.sideLength}
		area := sq.getArea()
		if area != test.area {
			t.Errorf("Expected square area %v to be %v", area, test.area)
		}
	}
}

func TestPrintArea(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	sq := square{sideLength: 1}
	area := sq.getArea()
	printArea(sq)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := fmt.Sprint("The shape area is: ", area, "\n")
	if string(out) != expected {
		t.Errorf("Expected %q, got %q", expected, out)
	}
}
