package dbfunc

import (
	"database/sql"
	"log"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

var (
	insertToDB   = "insert into school (name, address, status) value (?, ?, ?)"
	updateStatus = "update school set status = ?, ministry_id = ? where id = ?"
)

const (
	statusActive          = 1
	statusApproved        = 2
	statusMinistrySuccess = 3
	statusMinistryFailed  = 4
)

type StudentReq struct {
	File    multipart.File
	Name    string
	Address string
}

type StudentResp struct {
	Code    int      `json:"code,omitempty"`
	Data    DataResp `json:"data,omitempty"`
	Message string   `json:"message,omitempty"`
}

type DataResp struct {
	Name string `json:"student_name,omitempty"`
	ID   int64  `json:"id,omitempty"`
}

func dbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pass"
	dbName := "sql"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println("[dbOpen] Error to open database ")
		return nil, err
	}
	return db, nil
}

// Activate set status 'ACTIVE' on database
func Activate(student StudentReq) (int64, error) {
	db, err := dbConn()
	if err != nil {
		log.Println("[Activate] Error to open database ")
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare(insertToDB)
	if err != nil {
		log.Println("[Activate] Error to prepare query ")
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(student.Name, student.Address, statusActive)
	if err != nil {
		log.Println("[Activate] Error execute query ")
		return 0, err
	}
	documentID, err := result.LastInsertId()
	if err != nil {
		log.Println("[Activate] Error to get student id ")
		return 0, err
	}
	return documentID, nil
}

// Approval set status 'APPROVE' on database
func Approval(studentID int64, ministryID int64) error {
	db, err := dbConn()
	if err != nil {
		log.Println("[Approval] Error to open database ")
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare(updateStatus)
	if err != nil {
		log.Println("[Approval] Error to prepare query ")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(statusApproved, ministryID, studentID)
	if err != nil {
		log.Println("[Approval] Error execute query ")
		return err
	}
	return nil
}
