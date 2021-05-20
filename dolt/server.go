package main

import (
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	engine *gin.Engine
	logger *zap.Logger
}

func newServer(logger *zap.Logger) *Server {
	engine := gin.New()
	engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(logger, false))
	return &Server{
		engine: engine,
		logger: logger,
	}
}

func (s *Server) RegRoute(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	s.engine.Handle(httpMethod, relativePath, handlers...)
}
