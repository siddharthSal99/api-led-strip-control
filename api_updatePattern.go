package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func updatePattern(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)
	var newPattern PatternUpdateRequest

	resp["old-pattern"] = fmt.Sprintf("%v\n", currState.DisplayPattern)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newPattern)
	if err != nil {
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	formattedPattern, exists := PatternSet[newPattern.NewPattern]

	if exists {
		currState.DisplayPattern = formattedPattern
		resp["new-pattern"] = fmt.Sprintf("%v\n", newPattern.NewPattern)
	} else {
		fmt.Fprintf(w, "The pattern type you have specified does not exist")
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

	mutex.Unlock()
}
