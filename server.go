package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var PaletteToColors = map[string]ColorCombination{
	"Christmas":  {0x00007d00, 0x00555555, 0x007d0000},
	"Valentines": {0x006D196D, 0x006C0000, 0x005F4266},
	"Springtime": {0x00306C40, 0x001E6516, 0x005C601F},
	"Summer":     {0x001F605E, 0x001F4B60, 0x0062163C},
	"Fall":       {0x0060320E, 0x0062163C, 0x0060530E},
	"Diwali":     {0x00160E60, 0x005B0965, 0x00650909},
	"Halloween":  {0x00654909, 0x00550965, 0x00654909},
	"Icicle":     {0x00276970, 0x00274C70, 0x00273670},
	"Patriot":    {0x007d0000, 0x0000007d, 0x00555555},
	"Fireplace":  {0x007A6706, 0x007A3D00, 0x007A0029},
	"Fairy":      {0x0066661F, 0x00176F69, 0x006F6C17},
}

var PatternSet = map[string]Pattern{
	"Solid":             Solid,
	"Alternating":       Alternating,
	"SingleGradient":    SingleGradient,
	"DoubleAlternating": DoubleAlternating,
}

var ForcedStateSet = map[string]EnableState{
	"AlwaysOff": AlwaysOff,
	"AlwaysOn":  AlwaysOn,
	"Interval":  Interval,
}

var loc, _ = time.LoadLocation("EST")

var currState = LEDStripState{
	Colors: ColorCombination{
		Color1: 0x00007d00,
		Color2: 0x00555555,
		Color3: 0x0000007d,
	},
	Palette:        "Fairy",
	StartTime:      time.Date(0, 0, 0, 16, 0, 0, 0, loc),
	EndTime:        time.Date(0, 0, 0, 23, 0, 0, 0, loc),
	ForcedState:    Interval,
	DisplayPattern: Alternating,
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API for controlling LED strip")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	port := os.Getenv("PORT")

	// called by web app client
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/palette", updatePalette).Methods("PUT")
	router.HandleFunc("/endTime", updateEndTime).Methods("PUT")
	router.HandleFunc("/pattern", updatePattern).Methods("PUT")
	router.HandleFunc("/enable", updateForcedState).Methods("PUT")
	router.HandleFunc("/startTime", updateStartTime).Methods("PUT")

	router.HandleFunc("/colors", createColorCombination).Methods("POST")
	router.HandleFunc("/user/new", createNewUser).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")

	router.HandleFunc("/colors", getColors).Methods("GET")
	router.HandleFunc("/enabled", isEnabled).Methods("GET")
	router.HandleFunc("/endTime", getEndTime).Methods("GET")
	router.HandleFunc("/pattern", getPattern).Methods("GET")
	router.HandleFunc("/state", getFullState).Methods("GET")
	router.HandleFunc("/startTime", getStartTime).Methods("GET")
	router.HandleFunc("/palettes/all", getAllPalettes).Methods("GET")
	router.HandleFunc("/patterns/all", getAllPatterns).Methods("GET")
	router.HandleFunc("/enable-states/all", getAllEnableStates).Methods("GET")

	// CORS Preflight Handler
	router.Methods("OPTIONS").HandlerFunc(corsHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))

}
