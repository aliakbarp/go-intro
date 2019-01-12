package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/aliakbarp/go-intro/ministry/db"
)

const (
	// ReqLimitSize this is max size
	ReqLimitSize = 2
	// MB this is megabyte
	MB = 1 << 20
)

type Data struct {
	Name string `json:"student_name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

type MinistryResponse struct {
	Code    string `json:"code,omitempty"`
	Data    Data   `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// Hello function for greeting hello and testing the service
func Hello(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Name: "Palku Lingtu",
		ID:   69,
	}
	err := httpResponse(&w, data)
	if err != nil {
		log.Println("Error writing a response")
	}
}

// StudentHandler to handler student info that coming
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	studentInfo, err := checkRequest(r)
	if err != nil {
		log.Println("[UploadHandler] Error in checking request validity ")
		return
	}
	id, err := db.SetToArchived(studentInfo)
	if err != nil {
		return
	}

	response := Data{
		Name: studentInfo.Name,
		ID:   id,
	}

	err = httpResponse(&w, response)
	if err != nil {
		log.Println("Error writing a response")
	}
}

func checkRequest(r *http.Request) (db.Req, error) {
	result := db.Req{}
	err := r.ParseMultipartForm(ReqLimitSize * MB)
	if err != nil {
		log.Println("[checkRequest] Request exceeds size limit")
		return result, err
	}
	file, _, err := r.FormFile("imageFile")
	if err != nil {
		log.Println("[checkRequest] Failed to parsing image")
		return result, err

	}
	defer file.Close()
	result.File = file
	result.Name = r.PostFormValue("studentName")
	result.Address = r.PostFormValue("studentAddress")
	if result.Name == "" || result.Address == "" {
		log.Println("[checkRequest] No student name and/or address found")
		return result, errors.New("There is no student name and/or address")
	}
	return result, nil
}
