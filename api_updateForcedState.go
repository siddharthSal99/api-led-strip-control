package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func updateForcedState(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)
	var newForcedState EnableUpdateRequest

	resp["old-forced-state"] = fmt.Sprintf("%v", currState.ForcedState)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newForcedState)
	if err != nil {
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	formattedForcedState, exists := ForcedStateSet[newForcedState.NewForcedState]

	if exists {
		currState.ForcedState = formattedForcedState
		resp["new-forced-state"] = fmt.Sprintf("%v", newForcedState.NewForcedState)
	} else {
		fmt.Fprintf(w, "The forced state you have specified does not exist")
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

	mutex.Unlock()
}
