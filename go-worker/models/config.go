package models

import (
	"encoding/json"
	"os"
)

type ConfigJson struct {
	CheckInTime    string `json:"check_in_time"`
	CheckOutTime   string `json:"check_out_time"`
	AbsenceQuota   int    `json:"absence_quota"`
	DailyWorkHours int    `json:"daily_work_hours"`
	WeekyWorkHours int    `json:"weekly_work_hours"`
}

var (
	json_path = "./config.json"
)

func (config_json ConfigJson) LoadConfig() (ConfigJson, error) {
	if _, err := os.Stat(json_path); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(json_path)

			// Ensure no creating file error
			if err != nil {
				return ConfigJson{}, err
			}

			defer file.Close()

			err = ConfigJson.WriteFile(ConfigJson{})

			// Ensure no error writting file
			if err != nil {
				return ConfigJson{}, err
			}
		} else {
			return ConfigJson{}, err
		}
	}

	plan, err := os.ReadFile(json_path)

	// Ensure no error when openning config file
	if err != nil {
		return ConfigJson{}, err
	}

	err = json.Unmarshal(plan, &config_json)

	return config_json, err
}

func (config_json ConfigJson) WriteFile() error {
	content, err := json.MarshalIndent(config_json, "", " ")

	// Ensure no error unparsing struct to json
	if err != nil {
		return err
	}

	err = os.WriteFile(json_path, content, 0644)

	return err
}
