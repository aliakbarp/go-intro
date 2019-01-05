package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/aliakbarp/student-database/school/dbfunc"
	"github.com/aliakbarp/student-database/school/service"
)

const (
	// ReqLimitSize this is max size
	ReqLimitSize = 2
	// MB this is megabyte
	MB = 1 << 20
)

// Hello function for greeting hello and testing the service
func Hello(w http.ResponseWriter, r *http.Request) {
	err := httpResponse(&w, dbfunc.StudentResp{
		Code:    1111,
		Message: "Testing API",
	})
	if err != nil {
		log.Println("[Hello] Error writing a response")
	}
}

// UploadHandler to handler image that coming
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	studentInfo, err := checkRequest(r)
	if err != nil {
		log.Println("[UploadHandler] Error in checking request validity ")
		err = httpResponse(&w, dbfunc.StudentResp{
			Code:    3333,
			Message: "Invalid request",
		})
		if err != nil {
			log.Println("[UploadHandler] Error writing a response")
			return
		}
	}

	response, err := InsertToDB(studentInfo)
	if err != nil {
		log.Println("[UploadHandler] Error inserting to database ")
		err = httpResponse(&w, dbfunc.StudentResp{
			Code:    4444,
			Message: "Failed to process data",
		})
		if err != nil {
			log.Println("[UploadHandler] Error writing a response")
			return
		}
	}

	err = httpResponse(&w, response)
	if err != nil {
		log.Println("[UploadHandler] Error writing a response")
		return
	}
}

// InsertToDB insert image info into database
func InsertToDB(student dbfunc.StudentReq) (dbfunc.StudentResp, error) {
	result := dbfunc.StudentResp{}
	studentID, err := dbfunc.Activate(student)
	if err != nil {
		log.Println("[InsertToDB] Error to activating student ")
		return result, err
	}
	ministryRes, err := service.SendDataToMinistry(student)
	if err != nil {
		log.Println("[InsertToDB] Error sending data to ministry ")
		return result, err
	}
	err = dbfunc.Approval(studentID, ministryRes.Data.ID)
	if err != nil {
		log.Println("[InsertToDB] Error setting approval status ")
		return result, err
	}
	result = dbfunc.StudentResp{
		Code: 1000,
		Data: dbfunc.DataResp{
			Name: ministryRes.Data.Name,
			ID:   studentID,
		},
		Message: "Success",
	}

	return result, nil
}

func checkRequest(r *http.Request) (dbfunc.StudentReq, error) {
	student := dbfunc.StudentReq{}
	err := r.ParseMultipartForm(ReqLimitSize * MB)
	if err != nil {
		log.Println("[checkRequest] Request exceeds size limit")
		return student, err
	}
	file, _, err := r.FormFile("image_file")
	if err != nil {
		log.Println("[checkRequest] Failed to parsing image")
		return student, err

	}
	defer file.Close()
	student.File = file
	student.Name = r.PostFormValue("student_name")
	student.Address = r.PostFormValue("student_address")
	if student.Name == "" || student.Address == "" {
		log.Println("[checkRequest] No student name and/or address found")
		return student, errors.New("There is no student name")
	}
	return student, nil
}
