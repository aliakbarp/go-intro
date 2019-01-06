package handler

import (
	"errors"
	"html/template"
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
	tmpl := template.Must(template.ParseFiles("upload.html"))
	tmpl.Execute(w, nil)
}

// UploadHandler to handler image that coming
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	type htmlResponse struct {
		Header string             `json:"header,omitempty"`
		Data   dbfunc.StudentResp `json:"html_data,omitempty"`
	}
	tmpl := template.Must(template.ParseFiles("view.html"))

	studentInfo, err := checkRequest(r)
	if err != nil {
		log.Println("[UploadHandler] Error in checking request validity ")
		tmpl.Execute(w, htmlResponse{
			Header: "Invalid request",
			Data: dbfunc.StudentResp{
				Code: 3333,
			},
		})
		return
	}

	response, err := InsertToDB(studentInfo)
	if err != nil {
		log.Println("[UploadHandler] Error inserting to database ")
		tmpl.Execute(w, htmlResponse{
			Header: "Error inserting data to database",
			Data: dbfunc.StudentResp{
				Code: 5555,
			},
		})
		return
	}
	tmpl.Execute(w, htmlResponse{
		Header: "Thank you",
		Data:   response,
	})
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
