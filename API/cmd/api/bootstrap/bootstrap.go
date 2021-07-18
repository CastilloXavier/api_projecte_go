package bootstrap

import (
	mooc2 "github.com/CastilloXavier/api_project_go/API/internal"
	creating2 "github.com/CastilloXavier/api_project_go/API/internal/creating"
	increasing2 "github.com/CastilloXavier/api_project_go/API/internal/increasing"
	inmemory2 "github.com/CastilloXavier/api_project_go/API/internal/platform/bus/inmemory"
	server2 "github.com/CastilloXavier/api_project_go/API/internal/platform/server"
	mysql2 "github.com/CastilloXavier/api_project_go/API/internal/platform/storage/mysql"
	"context"
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}
	var (
		commandBus = inmemory2.NewCommandBus()
		eventBus = inmemory2.NewEventBus()
	)

	courseRepository := mysql2.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating2.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing2.NewCourseCounterService()

	createCourseCommandHandler := creating2.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating2.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc2.CourseCreatedEventType,
		creating2.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)


	ctx, srv := server2.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	Host            string        `default:"localhost"`
	Port			uint		  `default:"8080"`
	ShutdownTimeout	time.Duration `default:"10s"`
	// Database configuration
	DbUser		string		  `default:"root"`
	DbPass		string		  `default:"Gs4569"`
	DbHost		string		  `default:"localhost"`
	DbPort		uint		  `default:"3306"`
	DbName		string		  `default:"codely"`
	DbTimeout	time.Duration `default:"5s"`
}
