package main

import (
	"net/http"
	"resource"
)

func main() {
	resource.Echo()
	resource.Invert()
	resource.Flatten()
	resource.Sum()
	resource.Multiply()
	http.ListenAndServe(":8080", nil)
}
