package service

import (
	"net/http"

	"github.com/bnkamalesh/webgo/v6"
	"github.com/bnkamalesh/webgo/v6/middleware/cors"
	"github.com/rs/zerolog/log"

	"time"
)

func NewHTTPServer() {
	cfg := &webgo.Config{
		Host:         Config.Host,
		Port:         Config.Port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 1 * time.Hour,
	}

	InitLoggerHandler()
	log.Debug().Msg("Init Router.")
	routes := InitRouter()
	router := webgo.NewRouter(cfg, routes...)
	router.UseOnSpecialHandlers(accessLog)
	router.Use(errLogger, accessLog, cors.CORS(nil))
	router.Start()
}

func InitRouter() []*webgo.Route {
	return []*webgo.Route{
		{
			Name:          "root",
			Method:        http.MethodGet,
			Pattern:       "/",
			Handlers:      []http.HandlerFunc{HomeHandler},
			TrailingSlash: true,
		},
	}
}
