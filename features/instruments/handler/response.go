package handler

import "project/features/instruments"

type AllDataResponse struct {
	Data map[string]instruments.Instruments `json:"instruments"`
}
