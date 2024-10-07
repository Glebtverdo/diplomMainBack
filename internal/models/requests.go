package models

type ObjectBody struct {
	Name    string     `json:"name"`
	Address string     `json:"address"`
	Coords  [2]float32 `json:"coords"`
}

type Request struct {
	Id     int        `json:"id"`
	Currency int `json:"currency"`
	Object ObjectBody `json:"object"`
}

type ValueObj struct {
	Value int `json:"value"`
}

type ElementObj struct {
	Distance ValueObj `json:"distance"`
	Duration ValueObj `json:"duration"`
	Status   string   `json:"status"`
}

type RowObj struct {
	Elements []ElementObj `json:"elements"`
}

type YandexApiMatrixRes struct {
	Rows []RowObj `json:"rows"`
}

type Paths struct {
	Path     []int `json:"path"`
	Distance int   `json:"distance"`
}

type polylineType struct {
	Points [][2]float32 `json:"points"`
}

type stepType struct {
	Length          float64      `json:"length"`
	Duration        float64      `json:"duration"`
	Mode            string       `json:"mode"`
	WaitingDuration float64      `json:"waiting_duration"`
	Polyline        polylineType `json:"polyline"`
}

type legType struct {
	Status string     `json:"status"`
	Steps  []stepType `json:"steps"`
}

type flagsType struct {
	HasTolls                 bool `json:"hasTolls"`
	HasNonTransactionalTolls bool `json:"hasNonTransactionalTolls"`
}

type routeType struct {
	Legs  []legType `json:"legs"`
	Flags flagsType `json:"flags"`
}

type YandexApiRes struct {
	Route       routeType `json:"route"`
	TrafficType string    `json:"traffic_type"`
}

type CoordsWithPosition struct {
	Coords string
	Index  int
}


