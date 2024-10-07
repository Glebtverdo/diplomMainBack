package models

type WorkDayAddress struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Coords   string `json:"coords"`
	IsActive bool   `json:"is_active"`
}
