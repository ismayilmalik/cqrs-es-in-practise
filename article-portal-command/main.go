package main

import (
	"database/sql"
	"os"

	"github.com/ismayilmalik/cqrs-es-in-practise/apputils"
)

func main() {

	var connectionString string
	var servicePort string

	if os.Getenv("RUN_IN_CONTAINER") {
		connectionString = os.Getenv("COMMAND_DB")
		servicePort := os.Getenv("SERVICE_PORT")
	} else {
		connectionString = "root:123456@tcp(db_usersvc:3306)/articlePortal?parseTime=true&multiStatements=true&allowNativePasswords=True"
		servicePort := ":5050"
	}

	db, err := sql.Open("mysql", connectionString)
	apputils.HandleErrorWithFatal(err)

	a := App {}
	a.Prepare(db)
	a.Run(servicePort)
}
