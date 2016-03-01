package gamadues

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

//TestMain - the main testing function
func TestMain(t *testing.M) {
	apiKey := flag.String("api-key", "", "The API key you want to serve")
	flag.Parse()
	if *apiKey == "" {
		fmt.Println("No API key. Aborting...")
		os.Exit(1)
	}
	os.Setenv("APIKEY", *apiKey)
	os.Exit(t.Run())
}
