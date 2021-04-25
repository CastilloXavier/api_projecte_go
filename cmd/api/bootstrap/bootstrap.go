package bootstrap

import (
	"api_project/internal/creating"
	"api_project/internal/platform/bus/inmemory"
	"api_project/internal/platform/server"
	"api_project/internal/platform/storage/mysql"
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPass = "Gs4569"
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
	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)

	createCourseCommandHandler := creating.NewCourseCommandHandler(courseRepository)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, creatingCoursesService)
	return srv.Run()
}