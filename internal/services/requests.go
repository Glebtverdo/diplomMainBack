package services

import (
	"diplomMainBack/internal/models"
	antsalg "diplomMainBack/internal/services/antsAlg"
	"diplomMainBack/internal/services/graph"
	"diplomMainBack/internal/store"
	"diplomMainBack/internal/utils"
	"fmt"
)

func GetRequesCount(tokenPair models.JwtTokenPair) (int, error) {
	requests, err := store.GetRequests(tokenPair.Access)
	if err != nil {
		return 0, err
	}
	return len(requests), err
}

func makeCoords(requests []models.Request, destinationCoords string) []string {
	coords := []string{destinationCoords}

	for _, request := range requests {
		coords = append(coords, fmt.Sprintf("%f,%f", request.Object.Coords[0], request.Object.Coords[1]))
	}
	return coords
}

func initMatrix(len int) [][]float64 {
	matrix := make([][]float64, len)
	for i := range matrix {
		matrix[i] = make([]float64, len)
		matrix[i][i] = -1
	}
	return matrix
}

func findInitialRow(matrix [][]float64) int {
	for rowIndex, row := range matrix {
		if utils.IndexOf(row, 0, 0) != -1 {
			return rowIndex
		}
	}
	return -1
}

func fillMatrix(matrix [][]float64, distances []float64, indexes []int) [][]float64 {
	for index, dist := range distances {
		matrix[indexes[index]][indexes[index+1]] = dist
	}
	return matrix
}

func someMock(str []string) ([]float64, error) {
	var dist []float64
	for i := 0; i < len(str)-1; i++ {
		dist = append(dist, float64(i+1))
	}
	return dist, nil
}

func findNextPoint(matrix [][]float64, index int) int {
	res := utils.IndexOf(matrix[index], 0, index)
	if res == -1 {
		return utils.IndexOf(matrix[index], 0, 0)
	}
	return res
}

func getDistancesMatrix(coords []string) ([][]float64, error) {
	matrix := initMatrix(len(coords))
	index := findInitialRow(matrix)
	for requestsCount := 0; requestsCount < len(coords)*(len(coords)-1); requestsCount += 50 {
		coordsPack := []string{coords[index]}
		indexesPack := []int{index}
		for i := 0; i < 50; i++ {
			nextPoint := findNextPoint(matrix, index)
			if nextPoint == -1 {
				break
			}
			coordsPack = append(coordsPack, coords[nextPoint])
			indexesPack = append(indexesPack, nextPoint)
			matrix[index][nextPoint] = 1
			index = nextPoint
		}
		distances, err := store.GetCoordsList(coordsPack)
		fmt.Println(len(distances), distances)
		if err != nil {
			return nil, err
		}
		matrix = fillMatrix(matrix, distances, indexesPack)
	}
	return matrix, nil
}

func sortRoute(route models.RouteObl, coords []string) models.RouteObl {
	newRoute := models.RouteObl{
		Duration: route.Duration,
		Points: []models.RouteOblPoint{{
			Id:       0,
			Duration: 0,
			Coords:   coords[0],
		}},
	}
	firstPointIndex := 0
	for pointIndex, point := range route.Points {
		if point.Id == 0 {
			firstPointIndex = pointIndex
		}
	}
	for i := firstPointIndex + 1; i < len(route.Points); i++ {
		newRoute.Points = append(newRoute.Points, models.RouteOblPoint{
			Coords:   coords[route.Points[i].Id],
			Duration: route.Points[i].Duration,
			Id:       route.Points[i].Id,
		})
	}
	for i := 1; i <= firstPointIndex; i++ {
		newRoute.Points = append(newRoute.Points, models.RouteOblPoint{
			Coords:   coords[route.Points[i].Id],
			Duration: route.Points[i].Duration,
			Id:       route.Points[i].Id,
		})
	}
	return newRoute
}

func setRequestsId(route models.RouteObl, requests []models.Request) models.RouteObl {
	for i := 1; i < len(route.Points)-1; i++ {
		route.Points[i].RequestId = requests[route.Points[i].Id-1].Id
	}
	return route
}

func GetRouts(getRoutesModel models.GetRoutesModel) ([]models.RouteObl, error) {
	requests, err := store.GetRequests(getRoutesModel.Access)
	if err != nil {
		return nil, err
	}
	coords := makeCoords(requests, getRoutesModel.DestinationCoords)
	matrix, err := getDistancesMatrix(coords)
	if err != nil {
		return nil, err
	}
	// fmt.Println(coords)
	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%f, ", col)
	// 	}
	// 	fmt.Println()
	// }
	graphs := graph.MakeGraphs(len(matrix), getRoutesModel.PointsCount+1)
	var routes []models.RouteObl
	// antsalg.GetShortestPathFromGraph(graphs[len(graphs)-1], matrix)
	for _, g := range graphs {
		route := antsalg.GetShortestPathFromGraph(g, matrix)
		if route.Duration < 240 {
			route = sortRoute(route, coords)
			route = setRequestsId(route, requests)
			routes = append(routes, route)
		}
	}

	return routes, nil
}
