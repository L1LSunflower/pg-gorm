package utils

import (
	"encoding/json"
)

type Response struct {
	Status		string
	Code		int
	Messages	string
	Result		interface{}
}

func ResponseWrapper(data interface{}, statusCode int, message string) ([]byte, error){
	resp := Response{
		Status:   "ok",
		Code:     statusCode,
		Messages: message,
		Result:   data,
	}
	result, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return result, nil
}