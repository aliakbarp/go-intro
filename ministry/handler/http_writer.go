package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code    string      `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func httpResponse(WAddreess *http.ResponseWriter, data interface{}) error {
	w := *WAddreess
	result := Response{
		Code:    "1000",
		Message: "Success",
	}
	result.Data = data
	resp, err := json.Marshal(result)
	if err != nil {
		log.Println("Error writing a response")
	}
	fmt.Printf("%+v", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}
