package main

import (
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

// Creates a map that assigns each user a vault
// A vault is represented as a map that assigns a string to a string
var vaults = map[string]map[string]string{}

// Had to create a struct to make the process
// of sending back http errors easier
// based on taste, this doesn't feel like the best solution
// still searching
type ErrorStruct struct{
  Status int `json:"status"`
  Code   string `json:"code"`
}


func main(){

  // Create the mux router
  router := mux.NewRouter()

  // Define the routes
  router.HandleFunc("/api/account", createUser).Methods("POST")
  router.HandleFunc("/api/vault/{email}", getVaultContent).Methods("GET")
  router.HandleFunc("/api/vault/{email}", addKey).Methods("POST")
  router.HandleFunc("/api/vault/{email}", updateKey).Methods("PUT")
  router.HandleFunc("/api/vault/{email}", deleteKey).Methods("DELETE")

  // Start the server
  log.Fatal(http.ListenAndServe(":8000", router))

}


