package helper

import "encoding/json"

func ToMap(data any) (map[string]interface{}, error) {
	roomJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var roomMap map[string]interface{}
	if err := json.Unmarshal(roomJSON, &roomMap); err != nil {
		return nil, err
	}
	return roomMap, nil
}
