package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createColorCombination(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	// w.WriteHeader(http.StatusInternalServerError)
	// w.Header().Set("Content-Type", "application/json")

	// resp := make(map[string]string)
	// resp["message"] = "This endpoint has not been implemented"

	// jsonResp, err := json.Marshal(resp)
	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// }

	// w.Write(jsonResp)
	enableCors(&w)

	var newColorCombination ColorCombinationCreateRequest

	resp := make(map[string]string)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Unable to read request body, please make sure the request is formatted correctly")
	}

	err = json.Unmarshal(reqBody, &newColorCombination)
	if err != nil {
		fmt.Fprintf(w, "Unable to Unmarshal json, please make sure the request is formatted correctly")
	}

	PaletteToColors[newColorCombination.Name] = ColorCombination{
		Color1: newColorCombination.Color1,
		Color2: newColorCombination.Color2,
		Color3: newColorCombination.Color3,
	}

	currState.Colors = PaletteToColors[newColorCombination.Name]
	resp["new-color-combination"] = fmt.Sprintf("%v\n", PaletteToColors[newColorCombination.Name])

	mutex.Unlock()
}
