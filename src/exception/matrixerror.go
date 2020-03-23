package exception

import (
	"errors"
	"fmt"
)

//to do - cannot understand how to throw errors

// CheckMatrixTypes - check if the matrix have all elements as int type
func CheckMatrixTypes(records [][]string) error {
	var typeAccepted string = "int"
	var err error
	for _, row := range records {
		for _, num := range row {
			var actualType string = fmt.Sprintf("%T", num)
			if typeAccepted != actualType {
				err = errors.New("matrix doesn't contain only int types")
				break
			}
		}
	}
	return err
}
