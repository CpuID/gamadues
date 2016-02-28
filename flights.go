package gamadues

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
func (gm *Gamadeus) InspirationSearch(inspiR *InspirationRequest) *InspirationResult {
	if inspiR.AggregationMode != "" {
		if stringInSlice(inspiR.AggregationMode, getArray(aggregationModes)) == false {
			return nil
		}
	}
	result := InspirationResult{}
	gm.makeRequestGet("flights/inspiration-search?origin=BOS&departure_date=2016-03-01--2016-03-15&duration=7--9&max_price=500", &result)
	return &result
}
