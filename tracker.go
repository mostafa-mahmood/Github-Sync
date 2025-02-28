package main

import (
	"encoding/json"
	"os"
	"time"
)

func IncrementTimerTracker() error {
	data, err := ReadTimerJson()
	if err != nil {
		return err
	}

	data.TrackedMinutes += 5
	data.TotalSessionMinutes += 5
	data.LastUpdate = time.Now().Format("02/01/2006 15:04:05")

	jsonBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config/timer.json", jsonBytes, 0644)
}

func ResetTimerTracker() error {
	data, err := ReadTimerJson()
	if err != nil {
		return err
	}

	data.TrackedMinutes = 0
	data.LastUpdate = time.Now().Format("02/01/2006 15:04:05")

	jsonBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config/timer.json", jsonBytes, 0644)
}
