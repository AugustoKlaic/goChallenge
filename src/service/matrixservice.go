package service

import (
	"exception"
	"fmt"
	"strconv"
	"strings"
)

// PrintMatrixAsString - Transform the matrix to string
func PrintMatrixAsString(records [][]string) string {
	if !exception.CheckMatrix(records) { // validate matrix
		panic("Not valid Matrix")
	}

	var response string // variable that store matrix as string
	for _, row := range records {
		response = fmt.Sprintf("%s %s\n", response, strings.Join(row, ",")) // take response and add next line
	} // Sprintf dont print the string

	return response
}

// InvertMatrix - Transpose matrix
func InvertMatrix(records [][]string) string {
	x1 := len(records[0]) // store lenght of col
	y1 := len(records)    // store lenght of row

	response := make([][]string, x1) // create empty matrix
	for i := range response {
		response[i] = make([]string, y1)
	}

	for i := 0; i < x1; i++ { // go through the lines
		for j := 0; j < y1; j++ { // go through the cols
			response[i][j] = records[j][i] // invert matrix
		}
	}

	return PrintMatrixAsString(response) // call the print method
}

// FlattenMatrix - Write matrix as 1 line string
func FlattenMatrix(records [][]string) string {
	if !exception.CheckMatrix(records) { // validate matrix
		panic("Not valid Matrix")
	}

	var response []string         // variable that store elements of matrix
	for _, row := range records { // go through the first layer of records
		for _, num := range row { // go through the second layer of records
			response = append(response, fmt.Sprintf("%s", num)) // add to response every number on the matrix
		}
	}
	return fmt.Sprintf("%s", strings.Join(response, ",")) // return the string formatted
}

// SumMatrix - return the sum of all elements in the matrix
func SumMatrix(records [][]string) int {
	if !exception.CheckMatrix(records) { // validate matrix
		panic("Not valid Matrix")
	}

	var sum int = 0               // start with zero to prevent errors
	for _, row := range records { // go through the first layer of records
		for _, num := range row { // go through the second layer of records
			currentNum, _ := strconv.Atoi(num) // convert the string number to a int
			sum += currentNum                  // sum all up
		}
	}
	return sum
}

// MultiplyMatrix - return the product of the multiplication
func MultiplyMatrix(records [][]string) int {
	if !exception.CheckMatrix(records) { // validate matrix
		panic("Not valid Matrix")
	}

	var multiply int = 1          // start with 1 to not zero the multiplication
	for _, row := range records { // go through the first layer of records
		for _, num := range row { // go through the second layer of records
			currentNum, _ := strconv.Atoi(num) // convert the string number to a int
			multiply = multiply * currentNum   // multiply all
		}
	}
	return multiply
}
