package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Arrival describes an individual arrival time prediction
type Arrival struct {
	Run string `json:"rn"`
	ArrivalTime string `json:"arrT"`
}

// TrainTracker contains a grouping of etas
type TrainTracker struct {
	Arrivals []Arrival `json:"eta"`
}

// Response of the TT API
type Response struct {
	TrainTracker TrainTracker `json:"ctatt"`
}

func main() {
	westernMapID := 41480
	response, err := http.Get(fmt.Sprintf("http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx?key=%s&mapid=%d&outputType=JSON", KEY, westernMapID))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("error fetching:%d", response.StatusCode)
	}

	rb, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := &Response{}
	err = json.Unmarshal(rb, r)

	for _, arrival := range r.TrainTracker.Arrivals {
		fmt.Printf("arrival: run %s at %s\n", arrival.Run, arrival.ArrivalTime)
	}
}
