package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Arrival describes an individual arrival time prediction
type Arrival struct {
	StationID            string `json:"staId"`
	StopID               string `json:"stpId"`
	StopName             string `json:"staNm"`
	StopDestination      string `json:"stpDe"`
	Run                  string `json:"rn"`
	Route                string `json:"rt"`
	DestinationStationID string `json:"destSt"`
	DestinationName      string `json:"destNm"`
	TrainDirection       string `json:"trDr"`
	PredictionDateTime   string `json:"string"`
	ArrivalTime          string `json:"arrT"`
	// "1" if approaching, "0" if due
	IsApproaching        string `json:"isApp"`
	// "1" if based on schedule, "0" if prediction
	IsScheduled          string `json:"isSch"`
	// "1" if delayed, "0" otherwise
	IsDelayed            string `json:"isDly"`
	// "1" if fault is detected, "0" otherwise
	HasFault             string `json:"isFlt"`
	Latitude             string `json:"lat"`
	Longitude            string `json:"long"`
	// In Degrees
	Heading              string `json:"heading"`
}

// TrainTracker contains a grouping of etas
type TrainTracker struct {
	ErrorCode string    `json:"errCd"`
	ErrorName string    `json:"errNm"`
	Arrivals  []Arrival `json:"eta"`
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
