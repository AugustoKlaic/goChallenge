package service

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// MatrixReceiver - This function will receive the file and return the matrix
func MatrixReceiver(r *http.Request, w http.ResponseWriter) [][]string {
	file, _, err := r.FormFile("file")            // here we get the file, ignore the second return and get the error if it has one
	fileErrorHandler(w, err)                      // function that will show the error
	defer file.Close()                            // defer will close the file when all is done
	records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
	fileErrorHandler(w, err)                      // function that will show the error
	return records
}

func fileErrorHandler(w http.ResponseWriter, err error) { // receive an error and make a response
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error, matrix has the wrong format. %s", err.Error())))
		return
	}
}
