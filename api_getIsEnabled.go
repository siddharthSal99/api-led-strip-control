package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func isEnabled(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	enableCors(&w)
	resp := make(map[string]interface{})

	intervalValid := false

	currTime := time.Now().In(loc)

	/*
		if currTime.Hour() is less than endTime.Hour() and greater than startTime.Hour() set to true
		else if currTime.Hour() == startTime.Hour()
			if currTime.Hour == endTime.hour()
				set to currTime.Minute is within minute interval
			else
				set to currTime.Minute() > startTime.Minute()
		else if currTime.Hour == endTime.hour()
				set to currTime.Minute() < endTime.Minute()

	*/

	if currTime.Hour() > currState.StartTime.Hour() && currTime.Hour() < currState.EndTime.Hour() {
		intervalValid = true
	} else if currTime.Hour() == currState.StartTime.Hour() {
		if currTime.Hour() == currState.EndTime.Hour() {
			intervalValid = currTime.Minute() >= currState.StartTime.Minute() && currTime.Minute() <= currState.EndTime.Minute()
		} else {
			intervalValid = currTime.Minute() >= currState.StartTime.Minute()
		}
	} else if currTime.Hour() == currState.EndTime.Hour() {
		intervalValid = currTime.Minute() <= currState.EndTime.Minute()
	}

	// hh := currTime.Hour()
	// jj := currState.StartTime.Hour()

	// hm := currTime.Minute()
	// jm := currState.StartTime.Minute()

	// log.Printf("%v %v %v %v", hh, jj, hm, jm)

	isEnabled := ((intervalValid && currState.ForcedState == Interval) || currState.ForcedState == AlwaysOn) && (currState.ForcedState != AlwaysOff)

	resp["time"] = fmt.Sprintf("%v", currTime)
	resp["enabled"] = isEnabled //fmt.Sprintf("%v", currTime)
	resp["forced-state"] = currState.ForcedState

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	// w.WriteHeader(http.StatusOK)
	mutex.Unlock()
}
