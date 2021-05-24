package server

import (
	"os"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var gServer = &Server{}

type Server struct {
	addr     string
	logger   *zap.Logger
	engine   *gin.Engine
	wd       string
	handlers []Handler
}

type Handler struct {
	httpMethod   string
	relativePath string
	handlers     []gin.HandlerFunc
}

func Init(addr, wd string, logger *zap.Logger) { gServer.Init(addr, wd, logger) }
func (s *Server) Init(addr, wd string, logger *zap.Logger) {
	var err error
	if logger == nil {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
	}

	if err = os.Chdir(wd); err != nil {
		panic(err)
	}

	s.addr = addr
	s.wd = wd
	s.logger = logger
	s.engine = gin.New()
	s.engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	s.engine.Use(ginzap.RecoveryWithZap(logger, false))
	s.logger.Info("set work dir", zap.String("wd", s.wd))

	for _, h := range s.handlers {
		s.engine.Handle(h.httpMethod, h.relativePath, h.handlers...)
	}
}

func (s *Server) RegRoute(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	s.handlers = append(s.handlers, Handler{
		httpMethod:   httpMethod,
		relativePath: relativePath,
		handlers:     handlers,
	})
}

func Serve() { gServer.Serve() }
func (s *Server) Serve() {
	s.logger.Info("server listen at " + s.addr)
	if err := s.engine.Run(s.addr); err != nil {
		s.logger.Error("server stop", zap.Error(err))
	}
}
