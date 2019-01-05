package db

import (
	"database/sql"
	"log"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

var (
	insertToDB = "insert into edu_ministry (name, address, status) value (?, ?, ?)"
)

const (
	statusArchived = 1
)

type StudentReq struct {
	File    multipart.File
	Name    string
	Address string
}

type StudentResp struct {
	Name    string `json:"student_name,omitempty"`
	ID      string `json:"student_id,omitempty"`
	Message string `json:"message,omitempty"`
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

type Req struct {
	Name    string
	File    multipart.File
	Address string
}

// SetToArchived set status 'ACTIVE' on database
func SetToArchived(student Req) (int64, error) {
	db, err := dbConn()
	if err != nil {
		log.Println("[SetToArchived] Error to open database ")
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare(insertToDB)
	if err != nil {
		log.Println("[Activate] Error to prepare query ")
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(student.Name, student.Address, statusArchived)
	if err != nil {
		log.Println("[Activate] Error execute query ")
		return 0, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		log.Println("[Activate] Error to get student_data_id ")
		return 0, err
	}
	return ID, nil
}
