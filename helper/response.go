package helper

func FormatResponse(message string, data any, code int) map[string]any {
	var response = map[string]any{}
	response["message"] = message
	response["code"] = code
	if data != nil {
		response["data"] = data
	}
	return response
}
