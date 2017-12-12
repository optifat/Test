package main

import(
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

type myItem struct{
  key string `json:"key"`
  id int `json:"id"`
}

var (
  itemsStore  = make(map[string]int)
  thiskey = 0
)


func myHandler(w http.ResponseWriter, r *http.Request){

    var itemToAdd myItem
    j, _ := json.Marshal(itemToAdd)

    itemsStore[string(j)] = thiskey
    thiskey += 1
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(string(thiskey)))
}

func main()  {
    r := mux.NewRouter()
    r.HandleFunc("/", myHandler).Methods("GET", "POST")
    log.Fatal(http.ListenAndServe(":8082", r))
}
