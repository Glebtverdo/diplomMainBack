package graph

func MakeGraphs(pointsCount int, graphSize int) [][]int {
	max := pointsCount - 1
	var (
		graphs        [][]int
		pointsCounter []int
		pointer       int = graphSize - 1
	)
	pointsCounter = []int{0}
	for i := 1; i < graphSize; i++ {
		pointsCounter = append(pointsCounter, i)
	}
	graphsCount := countGraphs(max, graphSize-1)
	for i := 0; i < graphsCount; i++ {
		if i != 0 {
			checkPointsCounter(&pointer, &pointsCounter, max)
		}
		var graph []int
		for i := 0; i < len(pointsCounter); i++ {
			graph = append(graph, pointsCounter[i])
		}
		graphs = append(graphs, graph)
	}

	return graphs
}

func factorial(num int) int {
	res := 1
	for i := 2; i <= num; i++ {
		res *= i
	}
	return res
}

func countGraphs(pointsCount int, graphSize int) int {
	return factorial(pointsCount) / (factorial(graphSize) * factorial(pointsCount-graphSize))
}

func checkPointsCounter(pointer *int, pointsCounter *[]int, max int) {
	if (*pointsCounter)[(*pointer)] == max-(len((*pointsCounter))-(*pointer)-1) {
		for (*pointsCounter)[(*pointer)] == max-(len((*pointsCounter))-(*pointer)-1) {
			(*pointer) -= 1
		}
		(*pointsCounter)[(*pointer)] += 1
		setPointsCounter((*pointer), pointsCounter)
		(*pointer) = len((*pointsCounter)) - 1
	} else {
		(*pointsCounter)[(*pointer)] += 1
	}
}

func setPointsCounter(startIndex int, pointsCounter *[]int) {
	for startIndex < len(*pointsCounter)-1 {
		(*pointsCounter)[startIndex+1] = (*pointsCounter)[startIndex] + 1
		startIndex++
	}
}
