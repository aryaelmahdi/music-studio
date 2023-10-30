package helper

import (
	"encoding/json"
	"errors"
)

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

func ExtractPrice(room map[string]interface{}) (int, error) {
	priceValue, ok := room["price"]
	if !ok {
		return 0, errors.New("price is missing in the room data")
	}

	priceFloat, ok := priceValue.(float64)
	if !ok {
		return 0, errors.New("price is not a valid number")
	}

	price := int(priceFloat)
	return price, nil
}
