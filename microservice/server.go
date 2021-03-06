package microservice

import (
	"fmt"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Live long and Prosper!\n")
}
func NewServer(host string, port string) *http.Server {
	addr := fmt.Sprintf("%s:%s", host, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
