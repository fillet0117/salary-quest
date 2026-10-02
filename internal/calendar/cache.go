package calendar

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(year int) ([]Calendar, error) {
	fileName := fmt.Sprintf("./data/calendar_%d.json", year)

	data, err := os.ReadFile(
		fileName,
	)
	if err != nil {
		return []Calendar{}, err
	}

	var days []Calendar

	if err := json.Unmarshal(data, &days); err != nil {
		return []Calendar{}, err
	}

	return days, nil
}

func WriteFile(days []Calendar, year int) error {
	data, err := json.MarshalIndent(days, "", "  ")
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("./data/calendar_%d.json", year)

	err = os.WriteFile(
		fileName,
		data,
		0644,
	)

	return err
}
