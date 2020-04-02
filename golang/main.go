package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Entered in init")
}

func main() {
	fmt.Println("Entered in main")
	initRouter()
}

func initRouter() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/demo", demo)

	fmt.Println(http.ListenAndServe(":8080", router))

}
func demo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Demo OK")
	json.NewEncoder(w).Encode("Demo Json OK")
	w.WriteHeader(http.StatusOK)
}
