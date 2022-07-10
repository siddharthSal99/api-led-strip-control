package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createNewUser(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Header().Set("Content-Type", "application/json")

	// resp := make(map[string]string)
	// resp["message"] = "This endpoint has not been implemented"

	// jsonResp, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// }

	// w.Write(jsonResp)
	enableCors(&w)

	var newUser CreateNewUserRequest

	resp := make(map[string]string)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newUser)
	if err != nil {
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	usernameToPassword[newUser.Username] = newUser.Password
	usernameToAuthLevel[newUser.Username] = AuthLevel(newUser.AuthLevel)
	resp["new-user"] = fmt.Sprintf("%s: user successfully created\n", newUser.Username)

	mutex.Unlock()
}
