package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Struct to capture vault input
type VaultInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// This function is used to get the current
// vault content for a particular user email
// email is passed as a url param
func getVaultContent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]

	// Make sure a vault does exist
	// if not vault found, respond accordingly
	_, hasVault := vaults[email]
	if hasVault == false {
		json.NewEncoder(w).Encode(ErrorStruct{404, "ACCOUNT_NOT_FOUND"})
		return
	}

	json.NewEncoder(w).Encode(vaults[email])

}

// Function addkey adds a new key value pair to the vault
// Vault must not already contain the key
func addKey(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]

	// Check if the vault does exist
	vault, hasVault := vaults[email]
	if hasVault == false {
		json.NewEncoder(w).Encode(ErrorStruct{404, "ACCOUNT_NOT_FOUND"})
		return
	}

	// Grab instructions from the user
	vaultData := VaultInput{}

	err := json.NewDecoder(r.Body).Decode(&vaultData)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorStruct{400, "BAD_REQUEST_ERROR"})
		return
	}

	// Make sure the key does not already exist
	// return an error if the key is found
	_, hasKey := vault[vaultData.Key]
	if hasKey == true {
		json.NewEncoder(w).Encode(ErrorStruct{403, "KEY_EXISTS"})
		return
	}

	// Add key value pair to the vault
	vault[vaultData.Key] = vaultData.Value

	// Return current state of the vault
	json.NewEncoder(w).Encode(vault)

}

// Function updateKey changes the value of an already existing key
// Vault must exists and key must already exist
// returns error otherwise
func updateKey(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]

	// Make sure the vault does exist
	vault, hasVault := vaults[email]
	if hasVault == false {
		json.NewEncoder(w).Encode(ErrorStruct{404, "ACCOUNT_NOT_FOUND"})
		return
	}

	// Grab instructions from the user
	vaultData := VaultInput{}

	err := json.NewDecoder(r.Body).Decode(&vaultData)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorStruct{400, "BAD_REQUEST_ERROR"})
		return
	}

	// Make sure the key is infact in the vault
	_, hasKey := vault[vaultData.Key]
	if hasKey == false {
		json.NewEncoder(w).Encode(ErrorStruct{403, "KEY_DOES_NOT_EXIST"})
		return
	}

	// Update that key value pair
	vault[vaultData.Key] = vaultData.Value

	// Return the current state of the vault
	json.NewEncoder(w).Encode(vault)

}

// Function deleteKey deletes a key value pair from the vault
// Returns an error if account not found or key doesn't exist
func deleteKey(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	email := mux.Vars(r)["email"]

	// Make sure the vault does exist
	vault, hasVault := vaults[email]
	if hasVault == false {
		json.NewEncoder(w).Encode(ErrorStruct{404, "ACCOUNT_NOT_FOUND"})
		return
	}

	// Grab the user post data to determine what to delete
	vaultData := VaultInput{}

	err := json.NewDecoder(r.Body).Decode(&vaultData)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorStruct{400, "BAD_REQUEST_ERROR"})
		return
	}

	// Check that the vault does infact have the key to be deleted
	// return error if it doesn't
	_, hasKey := vault[vaultData.Key]
	if hasKey == false {
		json.NewEncoder(w).Encode(ErrorStruct{403, "KEY_DOES_NOT_EXIST"})
		return
	}

	// Actually delete the key
	delete(vault, vaultData.Key)

	// Return the current state of the vault
	json.NewEncoder(w).Encode(vault)

}
