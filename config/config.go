package config

import (
	"math/rand"
)

// TrafficLightConfig a configuration holding values related to the duration of a light period of each color or combination of colors.
type TrafficLightConfig struct {
	Green     int32 `json:"greenLightDuration"`
	Yellow    int32 `json:"yellowLightDuration"`
	YellowRed int32 `json:"yellowRedLightDuration"`
	RedLower  int32 `json:"lowerIntervalBorder"`
	RedUpper  int32 `json:"upperIntervalBorder"`
}

// RandRed returns a random light period for the red light within the specified limits.
func (tc TrafficLightConfig) RandRed() int32 {
	return rand.Int31n(tc.RedUpper-tc.RedLower) + tc.RedLower
}

// IsValid checks if the current values of the configuration are valid.
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
