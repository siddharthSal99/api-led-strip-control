package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

func getAllEnableStates(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()

	enableCors(&w)

	resp := make(map[string][]string)
	resp["enable-states"] = []string{}

	for k := range ForcedStateSet {
		resp["enable-states"] = append(resp["enable-states"], k)
	}

	sort.Strings(resp["enable-states"])

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	mutex.Unlock()
}
