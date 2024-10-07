package models

type GetRoutesModel struct {
	Access            string `json:"access"`
	Refresh           string `json:"refresh"`
	DestinationCoords string `json:"destination_coords"`
	PointsCount       int    `json:"points_count"`
}

type RouteOblPoint struct {
	Duration  float64 `json:"duration"`
	Id        int     `json:"id"`
	Coords    string  `json:"coords"`
	RequestId int     `json:"request_id"`
}

type RouteObl struct {
	Duration float64         `json:"duration"`
	Points   []RouteOblPoint `json:"points"`
}

type GetRoutesResponce struct {
	Count  int        `json:"count"`
	Routes []RouteObl `json:"routes"`
}
