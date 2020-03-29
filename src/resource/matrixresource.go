package resource

import (
	"fmt"
	"net/http"
	"service"
)

// Echo - This endpoint will print the matrix as string
func Echo() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		response := service.PrintMatrixAsString(service.MatrixReceiver(r, w)) // Call fileHandler
		fmt.Fprint(w, response)                                               // Now we print the entire matrix as string
	})
}

// Invert - This endpoint will invert the matrix and print it
func Invert() {
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		response := service.InvertMatrix(service.MatrixReceiver(r, w)) // Call fileHandler and then the matrix function
		fmt.Fprint(w, response)                                        // Now we print the entire matrix as string
	})
}

// Flatten - This endpoint will write matrix as 1 line string
func Flatten() {
	http.HandleFunc("/flat", func(w http.ResponseWriter, r *http.Request) {
		response := service.FlattenMatrix(service.MatrixReceiver(r, w)) // Call fileHandler and then the matrix function
		fmt.Fprint(w, response)                                         // Now we print the entire matrix as string
	})
}

// Sum - sum all the elements of the matrix
func Sum() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		response := service.SumMatrix(service.MatrixReceiver(r, w)) // Call fileHandler and then the matrix function
		fmt.Fprint(w, response)                                     // Now we print the entire matrix as string
	})
}

// Multiply - multiply all the elements of the matrix
func Multiply() {
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		response := service.MultiplyMatrix(service.MatrixReceiver(r, w)) // Call fileHandler and then the matrix function
		fmt.Fprint(w, response)                                          // Now we print the entire matrix as string
	})
}
