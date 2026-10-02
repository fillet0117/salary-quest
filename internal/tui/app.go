package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const totalSalary float64 = 60000.0

type tickMsg time.Time

type model struct {
	workDays    int
	daySalary   float64
	monthSalary float64
	perSalary   float64
	isHoliday   bool
}

func NewModel(
	workDays int,
	alreadyWork int,
	daySalary float64,
	monthSalary float64,
	isHoliday bool,
) model {
	perSalary := totalSalary / float64(workDays*24*60*60)

	return model{
		workDays:    workDays,
		daySalary:   daySalary,
		monthSalary: monthSalary,
		perSalary:   perSalary,
		isHoliday:   isHoliday,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			return m, tickCmd()
		}

	case tickMsg:
		m.daySalary += m.perSalary * 10
		m.monthSalary += m.perSalary * 10

		return m, tickCmd()
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	s += fmt.Sprintf("今天已經賺: $%.2f", m.daySalary)

	s += "\n\n"

	s += fmt.Sprintf("本月已賺: $%.2f", m.monthSalary)

	s += "\n\n"

	s += "enter 開始"

	return s
}

func tickCmd() tea.Cmd {
	return tea.Tick(
		10*time.Second,
		func(t time.Time) tea.Msg {
			return tickMsg(t)
		},
	)
}
