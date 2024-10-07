package store

import (
	"bytes"
	"diplomMainBack/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var mockJSON = `{"rows":[{"elements":[{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":5847},"duration":{"value":780}},{"status":"OK","distance":{"value":5407},"duration":{"value":1194}},{"status":"OK","distance":{"value":3848},"duration":{"value":925}},{"status":"OK","distance":{"value":3745},"duration":{"value":763}},{"status":"OK","distance":{"value":5407},"duration":{"value":1194}},{"status":"OK","distance":{"value":7674},"duration":{"value":996}},{"status":"OK","distance":{"value":4582},"duration":{"value":947}},{"status":"OK","distance":{"value":4806},"duration":{"value":977}},{"status":"OK","distance":{"value":4582},"duration":{"value":947}}]},{"elements":[{"status":"OK","distance":{"value":2369},"duration":{"value":411}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":6600},"duration":{"value":811}},{"status":"OK","distance":{"value":2236},"duration":{"value":517}},{"status":"OK","distance":{"value":6824},"duration":{"value":761}},{"status":"OK","distance":{"value":6600},"duration":{"value":811}},{"status":"OK","distance":{"value":1850},"duration":{"value":298}},{"status":"OK","distance":{"value":8173},"duration":{"value":1031}},{"status":"OK","distance":{"value":8122},"duration":{"value":882}},{"status":"OK","distance":{"value":8173},"duration":{"value":1031}}]},{"elements":[{"status":"OK","distance":{"value":2369},"duration":{"value":656}},{"status":"OK","distance":{"value":5496},"duration":{"value":982}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":2914},"duration":{"value":782}},{"status":"OK","distance":{"value":3382},"duration":{"value":784}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":7324},"duration":{"value":1197}},{"status":"OK","distance":{"value":3683},"duration":{"value":618}},{"status":"OK","distance":{"value":3907},"duration":{"value":647}},{"status":"OK","distance":{"value":3683},"duration":{"value":618}}]},{"elements":[{"status":"OK","distance":{"value":2460},"duration":{"value":386}},{"status":"OK","distance":{"value":1322},"duration":{"value":278}},{"status":"OK","distance":{"value":3526},"duration":{"value":717}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":4025},"duration":{"value":532}},{"status":"OK","distance":{"value":3526},"duration":{"value":717}},{"status":"OK","distance":{"value":3150},"duration":{"value":496}},{"status":"OK","distance":{"value":5551},"duration":{"value":936}},{"status":"OK","distance":{"value":5048},"duration":{"value":789}},{"status":"OK","distance":{"value":5551},"duration":{"value":936}}]},{"elements":[{"status":"OK","distance":{"value":2684},"duration":{"value":624}},{"status":"OK","distance":{"value":3244},"duration":{"value":443}},{"status":"OK","distance":{"value":5721},"duration":{"value":1130}},{"status":"OK","distance":{"value":5231},"duration":{"value":855}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":5721},"duration":{"value":1130}},{"status":"OK","distance":{"value":5071},"duration":{"value":660}},{"status":"OK","distance":{"value":8912},"duration":{"value":1274}},{"status":"OK","distance":{"value":7243},"duration":{"value":1201}},{"status":"OK","distance":{"value":8912},"duration":{"value":1274}}]},{"elements":[{"status":"OK","distance":{"value":2369},"duration":{"value":656}},{"status":"OK","distance":{"value":5496},"duration":{"value":982}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":2914},"duration":{"value":782}},{"status":"OK","distance":{"value":3382},"duration":{"value":784}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":7324},"duration":{"value":1197}},{"status":"OK","distance":{"value":3683},"duration":{"value":618}},{"status":"OK","distance":{"value":3907},"duration":{"value":647}},{"status":"OK","distance":{"value":3683},"duration":{"value":618}}]},{"elements":[{"status":"OK","distance":{"value":6613},"duration":{"value":764}},{"status":"OK","distance":{"value":4685},"duration":{"value":764}},{"status":"OK","distance":{"value":6440},"duration":{"value":988}},{"status":"OK","distance":{"value":6480},"duration":{"value":870}},{"status":"OK","distance":{"value":7736},"duration":{"value":1045}},{"status":"OK","distance":{"value":6440},"duration":{"value":988}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":8465},"duration":{"value":1207}},{"status":"OK","distance":{"value":7962},"duration":{"value":1060}},{"status":"OK","distance":{"value":8465},"duration":{"value":1207}}]},{"elements":[{"status":"OK","distance":{"value":4425},"duration":{"value":851}},{"status":"OK","distance":{"value":6708},"duration":{"value":878}},{"status":"OK","distance":{"value":2826},"duration":{"value":502}},{"status":"OK","distance":{"value":7713},"duration":{"value":1092}},{"status":"OK","distance":{"value":9411},"duration":{"value":1132}},{"status":"OK","distance":{"value":2826},"duration":{"value":502}},{"status":"OK","distance":{"value":6575},"duration":{"value":978}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":387},"duration":{"value":61}},{"status":"OK","distance":{"value":0},"duration":{"value":0}}]},{"elements":[{"status":"OK","distance":{"value":5223},"duration":{"value":1149}},{"status":"OK","distance":{"value":8167},"duration":{"value":948}},{"status":"OK","distance":{"value":3624},"duration":{"value":800}},{"status":"OK","distance":{"value":9172},"duration":{"value":1163}},{"status":"OK","distance":{"value":10027},"duration":{"value":1278}},{"status":"OK","distance":{"value":3624},"duration":{"value":800}},{"status":"OK","distance":{"value":9995},"duration":{"value":1167}},{"status":"OK","distance":{"value":2773},"duration":{"value":576}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":2773},"duration":{"value":576}}]},{"elements":[{"status":"OK","distance":{"value":4425},"duration":{"value":851}},{"status":"OK","distance":{"value":6708},"duration":{"value":878}},{"status":"OK","distance":{"value":2826},"duration":{"value":502}},{"status":"OK","distance":{"value":7713},"duration":{"value":1092}},{"status":"OK","distance":{"value":9411},"duration":{"value":1132}},{"status":"OK","distance":{"value":2826},"duration":{"value":502}},{"status":"OK","distance":{"value":6575},"duration":{"value":978}},{"status":"OK","distance":{"value":0},"duration":{"value":0}},{"status":"OK","distance":{"value":387},"duration":{"value":61}},{"status":"OK","distance":{"value":0},"duration":{"value":0}}]}]}`

// 25.234369457896325,55.280222457968712|25.234369457896325,55.401544758961258&apikey=6e17a68d-b013-4797-8151-d6d7e7b135ff
func GetMatrix(coords []string) ([][]int, error) {
	url := "https://api.routing.yandex.net/v2/distancematrix?"
	origins := "origins="
	destinations := "destinations="
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		return nil, fmt.Errorf("noApiKeyInEnv")
	}
	apiKey = fmt.Sprintf("apikey=%s", apiKey)
	for i, coordinate := range coords {
		fmt.Println(coordinate)
		if i != len(coords)-1 {
			origins = fmt.Sprintf("%s%s|", origins, coordinate)
			destinations = fmt.Sprintf("%s%s|", destinations, coordinate)
		} else {
			origins = fmt.Sprintf("%s%s", origins, coordinate)
			destinations = fmt.Sprintf("%s%s", destinations, coordinate)
		}
	}
	url = fmt.Sprintf("%s%s&%s&%s", url, origins, destinations, apiKey)
	fmt.Println(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	var rows models.YandexApiMatrixRes
	err = json.NewDecoder(res.Body).Decode(&rows)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	if len(rows.Rows) == 0 {
		buf := bytes.NewBufferString(mockJSON)
		json.NewDecoder(buf).Decode(&rows)
	}
	var durationMatrix [][]int
	for _, rowInfo := range rows.Rows {
		var row []int
		for i, el := range rowInfo.Elements {
			if el.Distance.Value < 300 {
				fmt.Println(i-1, i, i+1)
			}
			row = append(row, el.Distance.Value)
		}
		durationMatrix = append(durationMatrix, row)
	}
	return durationMatrix, nil
}

func GetCoordsList(coords []string) ([]float64, error) {
	url := "https://api.routing.yandex.net/v2/route?"
	coordinatesString := "waypoints="
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		return nil, fmt.Errorf("noApiKeyInEnv")
	}
	apiKey = fmt.Sprintf("apikey=%s", apiKey)
	for i, coordinate := range coords {
		if i != len(coords)-1 {
			coordinatesString = fmt.Sprintf("%s%s|", coordinatesString, coordinate)
		} else {
			coordinatesString = fmt.Sprintf("%s%s", coordinatesString, coordinate)
		}
	}
	url = fmt.Sprintf("%s%s&%s", url, coordinatesString, apiKey)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := httpClient.Do(request)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	rows := models.YandexApiRes{}
	// buf := bytes.NewBufferString(mockRequest)
	json.NewDecoder(res.Body).Decode(&rows)
	var durations []float64
	var distances []float64
	for index, leg := range rows.Route.Legs {
		durations = append(durations, 0)
		distances = append(distances, 0)
		for _, step := range leg.Steps {
			durations[index] += step.Duration
			distances[index] += step.Length
		}
		if durations[index] < 300 {
			fmt.Println(durations[index], distances[index])
		}
	}
	defer res.Body.Close()
	return durations, nil
}
