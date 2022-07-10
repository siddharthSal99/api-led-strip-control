package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func updateStartTime(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)
	var newStartTime StartTimeUpdateRequest
	isErr := false

	resp["old-time"] = fmt.Sprintf("%v\n", currState.StartTime)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
		// w.WriteHeader(http.StatusBadRequest)
	}

	err = json.Unmarshal(reqBody, &newStartTime)
	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
		// w.WriteHeader(http.StatusBadRequest)
	}

	newTime := newStartTime.NewStartTime
	formattedTime, err := time.Parse(layoutTime, newTime)
	resp["new-time"] = fmt.Sprintf("%v\n", newTime)

	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to convert time to HH:MM:SS format, please make sure the request is formatted correctly")
	}

	if !isErr {
		currState.StartTime = formattedTime
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	mutex.Unlock()
}
