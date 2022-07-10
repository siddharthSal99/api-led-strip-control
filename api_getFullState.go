package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getFullState(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]interface{})

	intervalValid := false

	if hourIntervalValid := time.Now().Hour() > currState.StartTime.Hour() && time.Now().Hour() < currState.EndTime.Hour(); hourIntervalValid {
		intervalValid = true
	} else if (time.Now().Hour() == currState.StartTime.Hour()) || (time.Now().Hour() == currState.EndTime.Hour()) {
		if minuteIntervalValid := time.Now().Minute() >= currState.StartTime.Minute() && time.Now().Minute() <= currState.EndTime.Minute(); minuteIntervalValid {
			intervalValid = true
		}
	}

	isEnabled := ((intervalValid && currState.ForcedState == Interval) || currState.ForcedState == AlwaysOn) && (currState.ForcedState != AlwaysOff)

	resp["time"] = fmt.Sprintf("%v", time.Now())
	resp["enabled"] = isEnabled //fmt.Sprintf("%v", time.Now())
	resp["start-time"] = fmt.Sprintf("%s", currState.StartTime.Format(layoutTime))
	resp["colors"] = currState.Colors
	resp["pattern"] = fmt.Sprintf("%s", currState.DisplayPattern)
	resp["end-time"] = fmt.Sprintf("%s", currState.EndTime.Format(layoutTime))
	resp["palette"] = currState.Palette
	resp["forced-state"] = currState.ForcedState

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	// w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}