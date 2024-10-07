package models

type Ant struct {
	Checked  map[int]bool `json:"checked"`
	Path     []int        `json:"path"`
	Duration float64      `json:"duration"`
	IsElit   bool         `json:"is_elit"`
	Id       int          `json:"id"`
}
