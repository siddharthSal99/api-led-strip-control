package main

import (
	"sync"
	"time"
)

var counter int
var mutex = &sync.Mutex{}

type EnableState string
type CombinationEnum int
type Pattern string
type Color int
type AuthLevel string

type ColorCombination struct {
	Color1 Color
	Color2 Color
	Color3 Color
}

type LEDStripState struct {
	Colors         ColorCombination
	StartTime      time.Time
	EndTime        time.Time
	ForcedState    EnableState
	DisplayPattern Pattern
	Palette        string
}

const (
	AlwaysOff EnableState = "AlwaysOff"
	AlwaysOn              = "AlwaysOn"
	Interval              = "Interval"
)

const (
	NotLoggedIn AuthLevel = "NotLoggedIn"
	Basic                 = "Basic"
	Admin                 = "Admin"
)

const MaxLoginAge int = 24 // login provides access to the app for 1 day before login expiry

const (
	DefaultColorCombo CombinationEnum = iota
	Christmas
	Valentines
	Springtime
	Summer
	Fall
	Diwali
	Halloween
	Icicle
	Patriot
	Fireplace
	Fairy
)

const (
	DefaultPattern    Pattern = "Default"
	Solid             Pattern = "Solid"
	Alternating       Pattern = "Alternating"
	SingleGradient    Pattern = "SingleGradient"
	DoubleGradient    Pattern = "DoubleGradient"
	DoubleAlternating Pattern = "DoubleAlternating"
)

const (
	layoutTime = "15:04:05"
)

type StartTimeUpdateRequest struct {
	NewStartTime string `json:"new-start-time"` //format: HH:MM:SS
}

type EndTimeUpdateRequest struct {
	NewEndTime string `json:"new-end-time"`
}

type PaletteUpdateRequest struct {
	NewPalette string `json:"new-palette"`
}

type PatternUpdateRequest struct {
	NewPattern string `json:"new-pattern"`
}

type EnableUpdateRequest struct {
	NewForcedState string `json:"new-forced-state"`
}

type ColorCombinationCreateRequest struct {
	Name   string `json:"name"`
	Color1 Color  `json:"color-1"`
	Color2 Color  `json:"color-2"`
	Color3 Color  `json:"color-3"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateNewUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AuthLevel string `json:"auth-level"`
}
