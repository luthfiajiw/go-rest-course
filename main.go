package main

import "net/http"

func main() {
	api := &api{Addr: ":8000"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.Addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	server.ListenAndServe()
}
