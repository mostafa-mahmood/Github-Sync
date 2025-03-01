package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Functinos to handel config directory

type TimerJson struct {
	TrackedMinutes      int    `json:"tracked_minutes"`
	TotalSessionMinutes int    `json:"total_session_minutes"`
	NumberOfCommits     int    `json:"number_of_commits"`
	LastUpdate          string `json:"last_update"`
}

type ConfigJson struct {
	GithubPAT       string `json:"Github_Personal_Access_Token"`
	Activity        string `json:"Activity"`
	CommitFrequency int    `json:"Commit_Frequency"`
}

func CreateConfigDirectories() {
	if !Exists("config") {
		os.MkdirAll("config", os.ModePerm)
	}

	if !Exists("repo") {
		os.MkdirAll("repo", os.ModePerm)
	}
}

func CreateConfigFiles() error {
	if !Exists("config/timer.json") {
		_, err := os.Create("config/timer.json")
		if err != nil {
			return fmt.Errorf("error creating timer.json: %v", err)
		}
		err = WriteTimerJson()
		if err != nil {
			return err
		}
	}

	if !Exists("config/config.json") {
		_, err := os.Create("config/config.json")
		if err != nil {
			return fmt.Errorf("error creating config.json: %v", err)
		}
		err = WriteConfigJson()
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteTimerJson() error {
	data := TimerJson{
		TrackedMinutes:      0,
		TotalSessionMinutes: 0,
		NumberOfCommits:     0,
		LastUpdate:          "Undefined",
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile("config/timer.json", jsonBytes, 0644)
}

func WriteConfigJson() error {
	data := ConfigJson{
		GithubPAT:       "Undefined",
		Activity:        "Undefined",
		CommitFrequency: 0,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("config/config.json", jsonBytes, 0644)
}

func ReadTimerJson() (TimerJson, error) {
	var data TimerJson

	jsonBytes, err := os.ReadFile("config/timer.json")
	if err != nil {
		return data, fmt.Errorf("error reading timer.json: %v", err)
	}

	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return data, fmt.Errorf("error parsing timer.json: %v", err)
	}

	return data, nil
}

func ReadConfigJson() (ConfigJson, error) {
	var data ConfigJson

	jsonBytes, err := os.ReadFile("config/config.json")
	if err != nil {
		return data, fmt.Errorf("error reading config.json: %v", err)
	}

	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return data, fmt.Errorf("error parsing config.json: %v", err)
	}

	return data, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
