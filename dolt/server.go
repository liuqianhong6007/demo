package main

import (
	"os"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var gServer = &Server{
	engine: gin.New(),
}

type Server struct {
	addr    string
	logger  *zap.Logger
	engine  *gin.Engine
	workDir string
}

func (s *Server) Init(addr, workDir string, logger *zap.Logger) {
	var err error
	if logger == nil {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	}

	if err = os.Chdir(workDir); err != nil {
		panic(err)
	}

	s.addr = addr
	s.workDir = workDir
	s.logger = logger
	s.engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	s.engine.Use(ginzap.RecoveryWithZap(logger, false))
	s.logger.Info("set work dir", zap.String("workdir", s.workDir))
}

func (s *Server) RegRoute(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	s.engine.Handle(httpMethod, relativePath, handlers...)
}

func (s *Server) Serve() {
	s.logger.Info("server listen at " + s.addr)
	if err := s.engine.Run(s.addr); err != nil {
		s.logger.Error("server stop", zap.Error(err))
	}
}