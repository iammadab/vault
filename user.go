package main

import (
  "net/http"
  "encoding/json"
)

// Struct to capture user creation request
type CreateUserRequest struct{
  Email string `json:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request){

  // Set the reply as json
  // Hate that I have to do this all the time
  // There has to be a better way
  w.Header().Set("Content-Type", "application/json")

  // Parsing json and sending json feels more complex
  // because json is not native to go
  // hopefully, there are better ways to handle this
  userData := CreateUserRequest{}

  err := json.NewDecoder(r.Body).Decode(&userData)
  if err != nil{
    json.NewEncoder(w).Encode(ErrorStruct{400, "BAD_REQUEST_ERROR"})
    return
  }

  // If the user already has a vault
  // then they already have an account
  _, hasVault := vaults[userData.Email]
  if hasVault == true{
    err := ErrorStruct{403, "ACCOUNT_EXISTS"}
    json.NewEncoder(w).Encode(err)
    return
  }

  // Create a vault for the user
  vaults[userData.Email] = map[string]string{}

  // Return the current user vault content
  // which will just be an empty object
  json.NewEncoder(w).Encode(vaults[userData.Email])

}
