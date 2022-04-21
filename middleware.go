package service

import (
	"net/http"
	"time"

	"github.com/bnkamalesh/webgo/v6"
	"github.com/rs/zerolog/log"
)

// errLogger is a middleware which will log all errors returned/set by a handler
func errLogger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
	err := webgo.GetError(r)
	if err != nil {
		if webgo.ResponseStatus(w) > 499 {
			log.Error().Err(err).Msg("ErrorLog")
		}
	}
}

// AccessLog is a middleware which prints access log to stdout
func accessLog(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, req)
	end := time.Now()
	log.Info().Str("Method", req.Method).
		Str("URL", req.URL.String()).
		Str("Consume", end.Sub(start).String()).
		Int("Status", webgo.ResponseStatus(rw)).
		Msg("AccessLog")
}
