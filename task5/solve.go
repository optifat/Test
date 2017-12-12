package main

import(
    //"encoding/json"
    "net/http"
    //"fmt"
    //"sort"
    //"strconv"

    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request){
  return
}

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/", handler)
  router.HandleFunc("/", handler).Methods("GET")
  http.ListenAndServe(":8082", router)
}
