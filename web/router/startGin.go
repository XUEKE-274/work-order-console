package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

var startGin = fx.Invoke(func(lifecycle fx.Lifecycle, g *gin.Engine, logger *logrus.Logger) {
	// Http Server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 启动服务
			go func() {
				httpServer.ListenAndServe()
			}()
			logger.Info("##################################")
			logger.Info("### WorkServiceConsole Success ###")
			logger.Info("##################################")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
})
