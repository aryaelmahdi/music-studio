package helper

import "encoding/json"

func ToMap(data any) (map[string]interface{}, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}
	return dataMap, nil
}
