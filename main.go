package main

import "net/http"

type api struct {
	addr string
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /users"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST /users"))
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	server.ListenAndServe()
}
