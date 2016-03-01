package gamadues

import (
	"os"
	"strconv"
	"testing"
	"time"
)

//TestInspirationSearch - inspiration search module
func TestInspirationSearch(t *testing.T) {
	if os.Getenv("APIKEY") == "" {
		t.Log("Could not find an API key")
		t.Fail()
	}
	apiKey := os.Getenv("APIKEY")
	gama := NewClient(apiKey)
	req := gama.GetInspirationRequest()
	req.Origin = "BOS"

	now := time.Now()
	month := "01"
	if int(now.Month()) < 10 {
		month = "0" + strconv.Itoa(int(now.Month()))
	}
	day := "01"
	if now.Day() < 10 {
		day = "0" + strconv.Itoa(now.Day())
	}
	nowStr := strconv.Itoa(now.Year()) + "-" + month + "-" + day
	nowPlus3 := time.Hour * 24 * 3
	diff := now.Add(nowPlus3)
	diffMonth := "01"
	if int(diff.Month()) < 10 {
		diffMonth = "0" + strconv.Itoa(int(diff.Month()))
	}
	diffDate := "03"
	if diff.Day() < 10 {
		diffDate = "0" + strconv.Itoa(diff.Day())
	}
	nowPlus3Str := strconv.Itoa(diff.Year()) + "-" + diffMonth + "-" + diffDate

	req.DepartureDate = nowStr + "--" + nowPlus3Str
	req.Duration = "7--9"
	req.MaxPrice = 500.00

	res, err := gama.InspirationSearch(*req)
	if err != nil {
		t.Logf("%+v", res)
		t.Log(err.Error())
		t.Fail()
	}
	if len(res.Results) == 0 {
		t.Log("0 Result from API")
		t.Fail()
	}

	if res.Origin != req.Origin {
		t.Log("Origins do not match in request and response")
		t.Fail()
	}

	for i := 0; i < len(res.Results); i++ {
		price, _ := strconv.ParseFloat(res.Results[i].Price, 64)
		if price > 500.00 {
			t.Log("Returned price is more than what we asked for")
			t.Fail()
		}
	}
}
