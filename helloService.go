package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {

	port := 8080

	http.HandleFunc("/hello", helloHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port), nil))


}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World\n")
}