package utils

import (
	_ "embed"
	"encoding/json"
	"os"
)

type Settings struct {
	OsuPath string `json:"osupath"`
}

func (s *Settings) Load() error {
	c, err := os.ReadFile("./settings.json")
	if err != nil {
		return err
	}
	json.Unmarshal(c, s)
	return nil
}

func (s *Settings) Save() error {
	str, err := json.Marshal(s)
	if err != nil {
		return err
	}

	os.WriteFile("./settings.json", str, 0644)

	return nil
}
