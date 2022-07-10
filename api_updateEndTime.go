package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func updateEndTime(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)
	var newEndTime EndTimeUpdateRequest
	isErr := false

	resp["old-time"] = fmt.Sprintf("%v\n", currState.EndTime)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newEndTime)
	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	newTime := newEndTime.NewEndTime
	formattedTime, err := time.Parse(layoutTime, newTime)
	resp["new-time"] = fmt.Sprintf("%v\n", newTime)

	if err != nil {
		isErr = true
		fmt.Fprintf(w, "Unable to convert time to HH:MM:SS format, please make sure the request is formatted correctly")
	}

	if !isErr {
		currState.EndTime = formattedTime
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	mutex.Unlock()
}
