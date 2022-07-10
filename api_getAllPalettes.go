package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
)

func getAllPalettes(w http.ResponseWriter, r *http.Request) {
	//return the list of keys in the ComboNameToColors map - this is for the drop down menu
	//of color choices on the front end
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string][]string)
	resp["palettes"] = []string{}
	for k := range PaletteToColors {
		resp["palettes"] = append(resp["palettes"], k)
	}

	sort.Strings(resp["palettes"])
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

	mutex.Unlock()
}
