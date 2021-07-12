package server

import (
	courses2 "api_project/API/internal/platform/server/handler/courses"
	health2 "api_project/API/internal/platform/server/handler/health"
	command2 "api_project/API/kit/command"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddr string
	engine 	*gin.Engine

	shutdownTimeout time.Duration

	// deps
	commandBus command2.Bus
}


func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command2.Bus) (context.Context, Server) {
	srv := Server{
		engine: gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health2.CheckHandler())
	s.engine.POST("/courses", courses2.CreateHandler(s.commandBus))
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on ", s.httpAddr)

	srv := &http.Server {
		Addr: s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func()  {
		<-c
		cancel()
	}()

	return ctx
}