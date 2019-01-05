package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type response struct {
	Name    string `json:"student_name,omitempty"`
	Message string `json:"message,omitempty"`
}

// DoToMinistry hit to ministry service
func DoToMinistry(url string, param MinistryReq, result interface{}) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("imageFile", "Image file")
	if err != nil {
		return err
	}

	_, err = io.Copy(part, param.Document)
	if err != nil {
		return err
	}

	err = writer.WriteField("studentName", param.Name)
	if err != nil {
		return err
	}

	err = writer.WriteField("studentAddress", param.Address)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/student", body)
	if err != nil {
		return err
	}

	httpClient := http.Client{}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := httpClient.Do(req)
	if err != nil {
		log.Println("[DoMinistry] Error when hitting edu ministry")
		return err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[DoMinistry] Error when reading response")
		return err
	}

	err = json.Unmarshal(contents, result)
	if err != nil {
		log.Println("[DoMinistry] Error when unmarshalling")
		return err
	}
	return nil
}
