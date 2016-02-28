package gamadues

import "errors"

//InspirationRequest - Defining the request body of Inspiration Search
type InspirationRequest struct {
	Origin          string  `json:"origin"`
	Destination     string  `json:"destination"`
	DepartureDate   string  `json:"departure_date"`
	OneWay          bool    `json:"one_way"`
	Duration        string  `json:"duration"`
	Direct          bool    `json:"direct"`
	MaxPrice        float64 `json:"max_price"`
	AggregationMode string  `json:"aggregation_mode"`
}

//InspirationResult - Defining the result body of the inspiration return
type InspirationResult struct {
	Origin   string `json:"origin"`
	Currency string `json:"currency"`
	Message  string `json:"message"`
	Status   int64  `json:"status"`
	Results  []struct {
		Destination   string `json:"destination"`
		DepartureDate string `json:"departure_date"`
		ReturnDate    string `json:"return_date"`
		Price         string `json:"price"`
		Airline       string `json:"airline"`
	} `json:"results"`
}

var aggregationModes = "DESTINATION,COUNTRY,DAY,WEEK"

//GetInspirationRequest - Get the request structure
func (gm *Gamadeus) GetInspirationRequest() *InspirationRequest {
	m := InspirationRequest{}
	return &m
}

//InspirationSearch - The interface to Inspiration search
func (gm *Gamadeus) InspirationSearch(inspiR InspirationRequest) (*InspirationResult, error) {
	if testAPIKey(gm) == false {
		return nil, errors.New("No API key")
	}
	if inspiR.Origin == "" {
		return nil, errors.New("Origin is a required field")
	}
	if inspiR.AggregationMode != "" {
		if stringInSlice(inspiR.AggregationMode, getArray(aggregationModes)) == false {
			return nil, errors.New("Aggregation mode is not supported")
		}
	}
	endPointParams := modifyToCallURL(inspiR)
	result := InspirationResult{}
	err := gm.makeRequestGet("flights/inspiration-search?"+endPointParams, &result)
	if err != nil {
		return &result, err
	}
	return &result, nil
}
