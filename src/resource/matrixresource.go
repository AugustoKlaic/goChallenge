package resource

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"service"
)

func fileErrorHandler(w http.ResponseWriter, err error) { // receive an error and make a response
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error, matrix has the wrong format. %s", err.Error())))
		return
	}
}

// Echo - This endpoint will print the matrix as string
func Echo() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file") // here we get the file, ignore the second return and get the error if it has one
		fileErrorHandler(w, err)           // function that will show the error
		defer file.Close()                 // defer will close the file when all is done

		records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
		fileErrorHandler(w, err)                      // function that will show the error

		response := service.PrintMatrixAsString(records)
		fmt.Fprint(w, response) // Now we print the entire matrix as string
	})
}

// Invert - This endpoint will invert the matrix and print it
func Invert() {
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file") // here we get the file, ignore the second return and get the error if it has one
		fileErrorHandler(w, err)           // function that will show the error
		defer file.Close()                 // defer will close the file when all is done

		records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
		fileErrorHandler(w, err)                      // function that will show the error

		response := service.InvertMatrix(records)
		fmt.Fprint(w, response) // Now we print the entire matrix as string
	})
}

// Flatten - This endpoint will write matrix as 1 line string
func Flatten() {
	http.HandleFunc("/flat", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file") // here we get the file, ignore the second return and get the error if it has one
		fileErrorHandler(w, err)           // function that will show the error
		defer file.Close()                 // defer will close the file when all is done

		records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
		fileErrorHandler(w, err)                      // function that will show the error

		response := service.FlattenMatrix(records)
		fmt.Fprint(w, response) // Now we print the entire matrix as string
	})
}

// Sum - sum all the elements of the matrix
func Sum() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file") // here we get the file, ignore the second return and get the error if it has one
		fileErrorHandler(w, err)           // function that will show the error
		defer file.Close()                 // defer will close the file when all is done

		records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
		fileErrorHandler(w, err)                      // function that will show the error

		response := service.SumMatrix(records)
		fmt.Fprint(w, response) // Now we print the entire matrix as string
	})
}

// Multiply - multiply all the elements of the matrix
func Multiply() {
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file") // here we get the file, ignore the second return and get the error if it has one
		fileErrorHandler(w, err)           // function that will show the error
		defer file.Close()                 // defer will close the file when all is done

		records, err := csv.NewReader(file).ReadAll() // read line per line of the matrix
		fileErrorHandler(w, err)                      // function that will show the error

		response := service.MultiplyMatrix(records)
		fmt.Fprint(w, response) // Now we print the entire matrix as string
	})
}
