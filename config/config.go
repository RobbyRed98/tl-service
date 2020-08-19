package config

import (
	"math/rand"
)

type TrafficLightConfig struct {
	Green     int32 `json:"greenLightDuration"`
	Yellow    int32 `json:"yellowLightDuration"`
	YellowRed int32 `json:"yellowRedLightDuration"`
	RedLower  int32 `json:"lowerIntervalBorder"`
	RedUpper  int32 `json:"upperIntervalBorder"`
}

func (tc TrafficLightConfig) RandRed() int32 {
	return rand.Int31n(tc.RedUpper-tc.RedLower) + tc.RedLower
}

func (tc TrafficLightConfig) IsValid() bool {
	if tc.Green <= 0 {
		return false
	} else if tc.Yellow <= 0 {
		return false
	} else if tc.YellowRed <= 0 {
		return false
	} else if tc.RedLower <= 0 {
		return false
	} else {
		return tc.RedLower <= tc.RedUpper
	}
}
