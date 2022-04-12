package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	tenant := os.Getenv("HELLOENV_TENANT")
	if tenant == "" || tenant == "TENANT_NOT_SET" {
		fmt.Println("HELLOENV_TENANT not set!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from "+tenant+"!")
	})

	fmt.Println("Go server is listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
