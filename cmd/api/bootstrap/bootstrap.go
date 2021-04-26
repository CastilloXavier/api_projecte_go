package bootstrap

import (
	"api_project/internal/creating"
	"api_project/internal/platform/bus/inmemory"
	"api_project/internal/platform/server"
	"api_project/internal/platform/storage/mysql"
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	host = "localhost"
	port = 8080
	shutdownTimeout = 10 * time.Second

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

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(),host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}