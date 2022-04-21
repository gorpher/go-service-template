package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gorpher/service"
	"github.com/rs/zerolog/log"
)

var version = "nil"
var hash = "nil"
var datetime = "nil"

func main() {
	v := flag.Bool("v", false, "查看版本号")
	flag.Parse()
	if *v {
		fmt.Printf("main-service has version %s built from %s on %s\n", version, hash, datetime)
		return
	}

	log.Debug().Msg("Init Environment.")
	initEnv()

	log.Debug().Msg("Init UUID.")
	if err := service.NewID(); err != nil {
		log.Fatal().Err(err).Caller().Send()
	}

	log.Debug().Msg("Init Database.")
	if err := service.InitDatabase(); err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Debug().Msg("Init Server.")
	service.NewHTTPServer()
}

func initEnv() {
	log.Debug().Msg("Init Config.")
	err := service.LoadDefaultYAMLConfig()
	if err != nil {
		log.Warn().Err(err).Send()
	}
	port := strings.TrimSpace(os.Getenv("HTTP_PORT"))
	if port != "" {
		service.Config.Port = port
	}
	log.Debug().Msg("Init Logger.")
	service.InitLogger(service.Config.LogLevel, service.Config.LogFile)
}
