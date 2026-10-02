package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const calendarUrl = "https://cdn.jsdelivr.net/gh/ruyut/TaiwanCalendar/data/"

func FetchCalendar(year int) ([]Calendar, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s%d.json", calendarUrl, year)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf(
			"fetch calendar: %w",
			err,
		)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"calendar api returned %s",
			resp.Status,
		)
	}

	var days []Calendar

	if err := json.NewDecoder(resp.Body).Decode(&days); err != nil {
		return nil, fmt.Errorf(
			"decode calendar: %w",
			err,
		)
	}

	return days, nil
}
