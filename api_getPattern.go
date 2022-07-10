package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getPattern(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)

	resp["pattern"] = fmt.Sprintf("%s", currState.DisplayPattern)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	// w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}
