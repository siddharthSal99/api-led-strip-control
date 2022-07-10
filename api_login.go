package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var usernameToPassword map[string]string = map[string]string{
	"ssalunkhe":          "9668",
	"sanjay.salunkhe":    "9668",
	"sampada.salunkhe":   "9668",
	"siddharth.salunkhe": "7437",
	"guest":              "0000",
	"SiddharthSal99":     "7437",
	"siddharthsal99":     "7437",
}

var usernameToAuthLevel map[string]AuthLevel = map[string]AuthLevel{
	"ssalunkhe":          Admin,
	"sanjay.salunkhe":    Admin,
	"sampada.salunkhe":   Admin,
	"siddharth.salunkhe": Admin,
	"guest":              Basic,
	"SiddharthSal99":     Admin,
	"siddharthsal99":     Admin,
}

var currentAuthLevel AuthLevel = NotLoggedIn
var currentUser = ""
var lastLogin time.Time = time.Now()

func login(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)

	timeSinceLastLogin := time.Now().Sub(lastLogin)

	if timeSinceLastLogin.Hours() > float64(MaxLoginAge) {
		currentAuthLevel = NotLoggedIn
		currentUser = ""
	}

	resp := make(map[string]interface{})
	loggedIn := currentAuthLevel != NotLoggedIn

	if !loggedIn {
		var loginRequest LoginRequest

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
		}

		err = json.Unmarshal(reqBody, &loginRequest)
		if err != nil {
			fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
		}

		password, usernameExists := usernameToPassword[loginRequest.Username]
		var success bool
		var message string
		if usernameExists && password == loginRequest.Password {
			currentAuthLevel = usernameToAuthLevel[loginRequest.Username]
			currentUser = loginRequest.Username
			success = true
			message = "Login Success"
			resp["username"] = currentUser
			lastLogin = time.Now()
		} else {
			success = false
			message = "Login Failed: Please verify username and password"
		}

		resp["success"] = fmt.Sprintf("%v", success)
		resp["message"] = fmt.Sprintf("%s", message)
		resp["authLevel"] = fmt.Sprintf("%v", currentAuthLevel)

	} else {
		resp["success"] = fmt.Sprintf("%v", true)
		resp["message"] = fmt.Sprintf("%s", "Login Success")
		resp["authLevel"] = fmt.Sprintf("%v", currentAuthLevel)
		resp["username"] = currentUser
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	// w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}

func isLoggedIn() bool {
	timeSinceLastLogin := time.Now().Sub(lastLogin)
	loggedIn := (currentAuthLevel != NotLoggedIn) && (timeSinceLastLogin.Hours() <= float64(MaxLoginAge))

	return loggedIn
}
