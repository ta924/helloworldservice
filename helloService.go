package main

import (
	"encoding/json"
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

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w,"Bad Request", http.StatusBadRequest)
		//ensure these are present or you will continue to process the request
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	//Marshal response and send success message with JSON
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}