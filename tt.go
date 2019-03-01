package tt

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

type resp struct {
	TrainTracker TrainTracker `json:"ctatt"`
}

// FetchArrivals returns all upcoming arrivals for a given stopID
// key is private CTA DLA API key
func FetchArrivals(key string, stopID int) (*TrainTracker, error) {
	response, err := http.Get(fmt.Sprintf("http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx?key=%s&mapid=%d&outputType=JSON", key, stopID))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching:%d", response.StatusCode)
	}

	rb, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	r := &resp{}
	err = json.Unmarshal(rb, r)

	return &r.TrainTracker, err
}
