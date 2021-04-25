package server

import (
	"api_project/internal/creating"
	"api_project/internal/platform/server/handler/courses"
	"api_project/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine 	*gin.Engine

	creatingCourseService creating.CourseSerice
}


func New(host string, port uint, creatingCourseService creating.CourseSerice) Server {
	srv := Server{
		engine: gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		creatingCourseService: creatingCourseService,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
}