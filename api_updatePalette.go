package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func updatePalette(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)

	resp := make(map[string]string)
	var newPalette PaletteUpdateRequest

	resp["old-colors"] = fmt.Sprintf("%v\n", currState.Colors)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newPalette)
	if err != nil {
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	newColors, exists := PaletteToColors[newPalette.NewPalette]
	if exists {
		currState.Colors = newColors
		currState.Palette = newPalette.NewPalette
		resp["new-colors"] = fmt.Sprintf("%v\n", newColors)
	} else {
		fmt.Fprintf(w, "The color palette name you have specified does not exist")
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	mutex.Unlock()
}
