package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getColors(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)

	resp := make(map[string]interface{})
	resp["colors"] = currState.Colors
	resp["palette"] = currState.Palette

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	mutex.Unlock()
}
