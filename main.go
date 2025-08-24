package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simplecontainer/smr/pkg/logger"
	"github.com/simplecontainer/traefik-provider/pkg/api"
	"github.com/simplecontainer/traefik-provider/pkg/configuration"
	"net/http"
	"os"
	"time"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	logger.Log = logger.NewLogger(logLevel, []string{"stdout"}, []string{"stderr"})

	config := configuration.New()
	api := api.New(config)

	router := gin.New()
	router.Use(gin.Recovery())

	go func() {
		err := api.Provider.Watch()
		if err != nil {
			panic(err)
		}
	}()

	provider := router.Group("/")
	{
		provider.GET("", api.HandleGetConfiguration)
		provider.GET("configuration", api.HandleGetConfiguration)

		provider.GET("http/routers", api.HandleGetHTTPRouters)
		provider.GET("http/services", api.HandleGetHTTPServices)
		provider.GET("http/middlewares", api.HandleGetHTTPMiddlewares)

		provider.GET("tcp/routers", api.HandleGetTCPRouters)
		provider.GET("tcp/services", api.HandleGetTCPServices)

		provider.GET("udp/routers", api.HandleGetUDPRouters)
		provider.GET("udp/services", api.HandleGetUDPServices)

		provider.GET("tls/certificates", api.HandleGetTLSCertificates)
		provider.GET("tls/options", api.HandleGetTLSOptions)

		provider.GET("health", api.HandleHealth)
		provider.POST("reload", api.HandleReload)
	}

	server := http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", api.Config.ProviderPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 0,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
