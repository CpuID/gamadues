package gamadues

import (
	"reflect"
	"strconv"
	"strings"
)

func getArray(input string) []string {
	if input == "" {
		return nil
	}
	return strings.Split(input, ",")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func testAPIKey(gm *Gamadeus) bool {
	if gm.APIKey.Key == "" {
		return false
	}
	return true
}

func addStr(input, name, value string) string {
	if input == "" {
		return name + "=" + value
	}
	return input + "&" + name + "=" + value
}

func modifyToCallURL(input interface{}) string {
	outPut := ""
	v := reflect.ValueOf(input)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			if v.Field(i).String() != "" {
				outPut = addStr(outPut, v.Type().Field(i).Tag.Get("json"), v.Field(i).String())
			}
		case reflect.Bool:
			if v.Field(i).Bool() == true || v.Field(i).Bool() == false {
				outPut = addStr(outPut, v.Type().Field(i).Tag.Get("json"), strconv.FormatBool(v.Field(i).Bool()))
			}
		case reflect.Float64:
			if v.Field(i).Float() > 0.00 {
				outPut = addStr(outPut, v.Type().Field(i).Tag.Get("json"), strconv.FormatFloat(v.Field(i).Float(), 'f', -1, 32))
			}
		case reflect.Int64:
			if v.Field(i).Int() > 0 {
				outPut = addStr(outPut, v.Type().Field(i).Tag.Get("json"), strconv.FormatInt(v.Field(i).Int(), 10))
			}
		}
	}
	return outPut
}
