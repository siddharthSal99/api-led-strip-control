package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getEndTime(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]string)

	resp["end-time"] = fmt.Sprintf("%s", currState.EndTime.Format(layoutTime))

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
	mutex.Unlock()
}
