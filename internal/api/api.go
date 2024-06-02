package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	fmt.Println("Init API")
	addr := ":8081"
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello server"))
	}).Methods("GET")
	http.ListenAndServe(addr, r)
}
