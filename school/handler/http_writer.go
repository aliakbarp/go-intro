package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func httpResponse(WAddreess *http.ResponseWriter, data interface{}) error {
	w := *WAddreess
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println("Error writing a response")
	}
	fmt.Printf("%+v", data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}
