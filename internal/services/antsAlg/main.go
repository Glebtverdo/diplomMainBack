package antsalg

import (
	"diplomMainBack/internal/models"
	"math"
	"math/rand"
)

func initAnts(graph []int) []models.Ant {
	var ants []models.Ant
	for i := 0; i < len(graph); i++ {
		ants = append(ants, models.Ant{
			Duration: 0,
			Path:     []int{graph[i]},
			Checked:  map[int]bool{graph[i]: true},
			IsElit:   false,
			Id:       i,
		})
		if (i+1)%5 == 0 {
			ants[i].IsElit = true
		}
	}
	return ants
}

func makeLabels(size int) [][]float64 {
	labels := make([][]float64, size)
	for rowIndex := 0; rowIndex < size; rowIndex++ {
		labels[rowIndex] = make([]float64, size)
		for colIndex := 0; colIndex < size; colIndex++ {
			if rowIndex == colIndex {
				labels[rowIndex][colIndex] = 0
			} else {
				labels[rowIndex][colIndex] = 1
			}
		}
	}
	return labels
}

func getPointToGo(ant models.Ant, labels [][]float64, graph []int, durationMatrix [][]float64) int {
	alf := float64(1)
	bet := float64(2)
	var hArr []float64
	hSum := float64(0)
	vertToGo := -1
	for _, point := range graph {
		H := float64(0)

		if (!ant.Checked[point] && durationMatrix[ant.Path[len(ant.Path)-1]][point] > 0) || (len(ant.Path) == len(graph) && point == ant.Path[0]) {
			N := 1 / durationMatrix[ant.Path[len(ant.Path)-1]][point]
			H = math.Pow(N, bet) * math.Pow(labels[ant.Path[len(ant.Path)-1]][point], alf)
			hSum += H
		}
		hArr = append(hArr, H)
	}
	if ant.IsElit {
		max := float64(0)
		for index, H := range hArr {
			if H != 0 && H > max {
				max = H
				vertToGo = graph[index]
			}
		}
	} else {
		num := rand.Intn(100) + 1
		roulette := float64(0)
		for index, H := range hArr {
			if H != 0 {
				roulette += H / hSum * 100
				if roulette >= float64(num) {
					vertToGo = graph[index]
					break
				}
			}
		}
	}
	return vertToGo
}

func removeFeramone(labels *[][]float64) {
	for rowIndex := 0; rowIndex < len(*labels); rowIndex++ {
		for colIndex := 0; colIndex < len(*labels); colIndex++ {
			(*labels)[rowIndex][colIndex] *= 0.5
		}
	}
}

func setAntsToDefault(ants []models.Ant, graph []int) []models.Ant {
	for antIndex, ant := range ants {
		for pathId, _ := range ant.Path {
			if pathId != graph[antIndex] {
				ants[antIndex].Checked[pathId] = false
			}
		}
		ants[antIndex].Path = []int{graph[antIndex]}
		ants[antIndex].Duration = 0
	}
	return ants
}

func GetShortestPathFromGraph(graph []int, durationMatrix [][]float64) models.RouteObl {
	ants := initAnts(graph)
	labels := makeLabels(len(durationMatrix))
	sameCount := 0
	Q := float64(10)
	min := float64(-1)
	minPaths := make([]int, len(graph)+1)
	for iter := 0; iter < 2000; iter++ {
		ants = setAntsToDefault(ants, graph)
		for step := 1; step <= len(graph); step++ {
			for i := 0; i < len(ants); i++ {
				pointToGo := getPointToGo(ants[i], labels, graph, durationMatrix)
				if pointToGo != -1 {
					ants[i].Checked[pointToGo] = true
					labels[ants[i].Path[len(ants[i].Path)-1]][pointToGo] += Q / durationMatrix[ants[i].Path[len(ants[i].Path)-1]][pointToGo]
					ants[i].Duration += durationMatrix[ants[i].Path[len(ants[i].Path)-1]][pointToGo]
					ants[i].Path = append(ants[i].Path, pointToGo)
				}
			}
			removeFeramone(&labels)
		}
		for _, ant := range ants {
			if (ant.Duration < min || min == -1) && ant.Path[0] == ant.Path[len(ant.Path)-1] {
				if min == ant.Duration {
					sameCount++
				} else {
					sameCount = 0
				}
				min = ant.Duration
				for i := 0; i < len(ant.Path); i++ {
					minPaths[i] = ant.Path[i]
				}
			}
		}
		if sameCount == 10 {
			break
		}
	}
	points := []models.RouteOblPoint{}
	for i, path := range minPaths {
		points = append(points, models.RouteOblPoint{
			Id: path,
		})
		if i == 0 {
			points[i].Duration = 0
		} else {
			points[i].Duration = durationMatrix[points[i-1].Id][path]
		}
	}
	return models.RouteObl{
		Duration: min,
		Points:   points,
	}
}
