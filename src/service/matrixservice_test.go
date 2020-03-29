package service

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	records := createMatrix()
	actual := PrintMatrixAsString(records)
	expected := " 8,16\n 9,18\n"

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestInvert(t *testing.T) {
	records := createMatrix()
	actual := InvertMatrix(records)
	expected := " 8,9\n 16,18\n"

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestFlatten(t *testing.T) {
	records := createMatrix()
	actual := FlattenMatrix(records)
	expected := "8,16,9,18"

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestSum(t *testing.T) {
	records := createMatrix()
	actual := SumMatrix(records)
	expected := 51

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestMultiply(t *testing.T) {
	records := createMatrix()
	actual := MultiplyMatrix(records)
	expected := 20736

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func createMatrix() [][]string {
	x1 := 2 // store lenght of col
	y1 := 2 // store lenght of row

	records := make([][]string, x1) // create empty matrix
	for i := range records {
		records[i] = make([]string, y1)
	}

	for i := 0; i < x1; i++ { // go through the lines
		for j := 0; j < y1; j++ { // go through the cols

			records[i][j] = fmt.Sprintf("%d", ((i+1)+7)*(j+1)) // fill up
		}
	}
	return records
}
