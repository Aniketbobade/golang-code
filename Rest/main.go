package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	
	// Start the server on port 8080
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}