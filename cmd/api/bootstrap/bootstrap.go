package bootstrap

import (
	"api_project/internal/platform/server"
	"api_project/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "gs3458"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)
func Run() error{
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepositroy(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}