package service

import (
	"mime/multipart"

	"github.com/aliakbarp/go-intro/school/dbfunc"
)

type MinistryReq struct {
	Name     string
	Document multipart.File
	Address  string
}

type Data struct {
	Name string `json:"student_name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

type MinistryResp struct {
	Code    string `json:"code,omitempty"`
	Data    Data   `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// SendDataToMinistry sent student data to education ministry database
func SendDataToMinistry(student dbfunc.StudentReq) (MinistryResp, error) {
	result := MinistryResp{}
	request := MinistryReq{
		Name:     student.Name,
		Document: student.File,
		Address:  student.Address,
	}
	err := DoToMinistry("http://localhost:8000/student", request, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
