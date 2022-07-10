package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

func getAllPatterns(w http.ResponseWriter, r *http.Request) {

	//return the list of keys in the PatternSet map - this is for the drop down menu
	//of patterns on the front end
	mutex.Lock()

	enableCors(&w)

	resp := make(map[string][]string)
	resp["patterns"] = []string{}

	for k := range PatternSet {
		resp["patterns"] = append(resp["patterns"], k)
	}

	sort.Strings(resp["patterns"])

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	mutex.Unlock()
}
