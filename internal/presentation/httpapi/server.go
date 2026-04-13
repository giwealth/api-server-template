package httpapi

import (
	"api-service-template/internal/infra"
	"fmt"
	"net"
	"net/http"

	"api-service-template/internal/option"
	mw "api-service-template/internal/presentation/httpapi/middlewares"
	"api-service-template/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Server http server
type Server struct {
	r *gin.Engine
	g *gin.RouterGroup
}

// NewServer 构造函数
func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	g := r.Group("/api")

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AddAllowHeaders("*")

	g.Use(cors.New(cfg))
	g.Use(mw.ValidateHeaders)
	g.Use(mw.RequestLogger)
	g.Use(gin.CustomRecoveryWithWriter(panicWriter{}, func(c *gin.Context, e interface{}) {
		if err, ok := e.(error); ok {
			mw.SaveRequestError(c, err)
			if rerr, ok := errors.Cause(err).(apiError); ok {
				c.JSON(rerr.Status(), rerr)
				return
			}
		}
		rerr := errInternal
		c.JSON(rerr.Status(), rerr)
	}))

	return &Server{r: r, g: g}
}

// Run 运行http api
func (s *Server) Run(opt *option.Options) (*http.Server, error) {
	if err := initDB(opt.GetDB()); err != nil {
		return nil, err
	}

	s.router(opt)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", opt.ListenPort),
		Handler: s.r,
	}

	ln, err := net.Listen("tcp4", fmt.Sprintf(":%v", opt.ListenPort))
	if err != nil {
		return nil, err
	}
	go func() {
		logger.DefaultEntry().Printf("start server, listen port %v", opt.ListenPort)
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			logger.DefaultEntry().Fatal("start server:", err)
		}
	}()

	return srv, nil
}

func initDB(db *gorm.DB) error {
	infra.Init(db)
	return nil
}

type panicWriter struct{}

func (w panicWriter) Write(p []byte) (n int, err error) {
	logrus.WithField("panic", string(p)).Error("request panic")
	return len(p), nil
}
