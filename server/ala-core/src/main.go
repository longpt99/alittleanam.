package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/longpt99/alittleanam/server/ala-core/src/utils"
)

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		utils.Write(res, req, nil)
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
