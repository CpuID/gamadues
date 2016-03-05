package gamadues

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//APICreds -  The datastructure to hold the API key
type APICreds struct {
	Key string
}

//Gamadeus - The main datastructure holding a pointer to APICreds
type Gamadeus struct {
	Version        string
	APIKey         *APICreds
	SandboxVersion string
	PrimaryURL     string
}

const version = "0.1"
const sandboxversion = "v1.2"
const baseurl = "http://api.sandbox.amadeus.com/"

//NewClient - Obtain a new client pointer to the API
func NewClient(key string) *Gamadeus {
	apiCred := APICreds{key}
	return &Gamadeus{version, &apiCred, sandboxversion, baseurl + sandboxversion}
}

//GetVersion - Get the version information of this library
func (gm *Gamadeus) GetVersion() string {
	return gm.Version
}

func (gm *Gamadeus) makeRequestGet(endPoint string, returnData interface{}) error {
	//This is a test comment. Thanks
	//Testing for the github webhook. ttttt
	callURL := gm.PrimaryURL + "/" + endPoint + "&apikey=" + gm.APIKey.Key
	//fmt.Println(callURL)
	res, err := http.Get(callURL)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	jsonDataFromHTTP, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonDataFromHTTP), &returnData)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("Non OK response received")
	}
	return nil
}
