package helper

import "encoding/json"

func StructToMapString(in interface{}) (map[string]interface{}, error) {

	var out map[string]interface{}
	byteData, _ := json.Marshal(in)
	err := json.Unmarshal(byteData, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}

func MapToJson(in map[string]interface{}) (string, error) {

	out, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(out), nil

}
