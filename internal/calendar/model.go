package calendar

type Calendar struct {
	Date        string `json:"date"`
	Week        string `json:"week"`
	IsHoliday   bool   `json:"isHoliday"`
	Description string `json:"description"`
}
