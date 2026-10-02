package main

import (
	"log"
	"salary-quest/internal/calendar"
	"salary-quest/internal/tui"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const totalSalary float64 = 60000.0

func main() {
	now := time.Now()

	// 取得行事曆
	days, err := GetCalendar(now.Year())
	if err != nil {
		log.Fatalf("取得行事曆失敗: %v", err)
	}

	var workDays int
	var isHoliday bool
	var alreadyWork int

	start := time.Date(
		now.Year(),
		now.Month(),
		1,
		0,
		0,
		0,
		0,
		now.Location(),
	)

	for _, day := range days {
		t, err := time.Parse("20060102", day.Date)
		if err != nil {
			continue
		}

		if t.Year() == now.Year() &&
			t.Month() == now.Month() {

			if t.Day() == now.Day() {
				isHoliday = day.IsHoliday
			}

			if !day.IsHoliday {
				workDays++
			}

			if t.Before(start) && !day.IsHoliday {
				alreadyWork++
			}
		}
	}

	// 每日工資
	daySalary := totalSalary / float64(workDays)

	monthSalary := float64(alreadyWork) * daySalary

	model := tui.NewModel(workDays, alreadyWork, daySalary, monthSalary, isHoliday)

	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		log.Fatalf("run err: %v", err)
	}
}

func GetCalendar(year int) ([]calendar.Calendar, error) {
	days, err := calendar.ReadFile(year)
	if err == nil && len(days) != 0 {
		return days, nil
	}

	fetchDays, err := calendar.FetchCalendar(year)
	if err != nil {
		return []calendar.Calendar{}, err
	}

	err = calendar.WriteFile(fetchDays, year)
	if err != nil {
		return []calendar.Calendar{}, err
	}

	return fetchDays, nil
}
