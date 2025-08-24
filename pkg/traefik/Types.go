package traefik

import (
	"github.com/gin-gonic/gin"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"go.uber.org/zap"
	"sync"
)

type TraefikConfiguration struct {
	HTTP *dynamic.HTTPConfiguration `json:"http,omitempty"`
	TCP  *dynamic.TCPConfiguration  `json:"tcp,omitempty"`
	UDP  *dynamic.UDPConfiguration  `json:"udp,omitempty"`
	TLS  *dynamic.TLSConfiguration  `json:"tls,omitempty"`
}

type ConfigurationManager struct {
	config *dynamic.Configuration
	mu     sync.RWMutex
	logger *zap.Logger
}

type Server struct {
	configManager *ConfigurationManager
	router        *gin.Engine
	logger        *zap.Logger
	port          string
}
