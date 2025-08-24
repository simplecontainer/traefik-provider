package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (a *Api) HandleGetConfiguration(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config)
}

func (a *Api) HandleGetHTTPRouters(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.HTTP.Routers)
}

func (a *Api) HandleGetHTTPServices(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.HTTP.Services)
}

func (a *Api) HandleGetHTTPMiddlewares(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.HTTP.Middlewares)
}

func (a *Api) HandleGetTCPRouters(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.TCP.Routers)
}

func (a *Api) HandleGetTCPServices(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.TCP.Services)
}

func (a *Api) HandleGetUDPRouters(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.UDP.Routers)
}

func (a *Api) HandleGetUDPServices(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.UDP.Services)
}

func (a *Api) HandleGetTLSCertificates(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.TLS.Certificates)
}

func (a *Api) HandleGetTLSOptions(c *gin.Context) {
	c.JSON(http.StatusOK, a.Provider.Config.TLS.Options)
}

func (a *Api) HandleHealth(c *gin.Context) {
	health := gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, health)
}
