package util

import "encoding/json"

func ToJSON(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(json)
}
