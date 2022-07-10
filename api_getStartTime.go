package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getStartTime(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)

	// currEndTime, err := time.Parse("HH:MM:SS", currState.EndTime.String())
	resp["start-time"] = fmt.Sprintf("%s", currState.StartTime.Format(layoutTime))

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	// w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}
