package obtainerDb

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

var creds map[string]string = map[string]string{
	"username": "SCOTT",
	"password": "Tiger777",
	"host":     "localhost",
	"port":     "1521",
	"sid":      "ORCLPDB1",
}

func GetDbConnection() *sql.DB {

	connection := "oracle://" + creds["username"] + ":" + creds["password"] + "@" + creds["host"] + ":" + creds["port"] + "/" + creds["sid"]
	db, error := sql.Open("oracle", connection)
	if error != nil {
		panic(fmt.Errorf("Database can not be open, error - %w", error))
	}
	err := db.Ping()
	if err != nil {
		panic(fmt.Errorf("Error pinging db - %w", err))
	}
	return db
}
