package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bnkamalesh/webgo/v6"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const IsoZonedDateTime = "2006-01-02 15:04:05"

// nolint
func init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: IsoZonedDateTime,
	}).With().Caller().Logger()
}

// InitLogger 初始化logger.
func InitLogger(level, path string) {
	switch level {
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "empty":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// 设置时间格式
	zerolog.TimeFieldFormat = IsoZonedDateTime

	// 设置日志输出方式
	var (
		consoleWriter io.Writer
		logFile       *os.File
		err           error
	)
	// 设置日志文件
	if path != "" {
		logFile, err = initLoggerFile(path)
		if err != nil {
			log.Error().Err(err).Msg("初始化日志文件失败")
		}
	}
	consoleWriter = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: IsoZonedDateTime,
	}
	if logFile != nil {
		log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter, logFile)).With().Timestamp().Logger()
	} else {
		log.Logger = log.Output(consoleWriter)
	}
	if level == "debug" {
		log.Logger = log.Logger.With().Caller().Logger()
	}
}

// initLoggerFile 初始化日志文件
func initLoggerFile(path string) (logFile *os.File, err error) {
	var location string
	if !filepath.IsAbs(path) {
		location, err = os.Executable()
		if err != nil {
			log.Error().Err(err).Msg("获取当前执行文件路径失败")
			return nil, err
		}
		path = filepath.Join(filepath.Dir(location), path)
	}
	log.Debug().Str("log_path", path).Msg("Init Logger")
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, os.ModePerm) // nolint
}

type loggerHandler struct {
}

func (l loggerHandler) Debug(data ...interface{}) {
	log.Debug().Msg(fmt.Sprint(data...))
}

func (l loggerHandler) Info(data ...interface{}) {
	log.Info().Msg(fmt.Sprint(data...))
}

func (l loggerHandler) Warn(data ...interface{}) {
	log.Warn().Msg(fmt.Sprint(data...))
}

func (l loggerHandler) Error(data ...interface{}) {
	log.Error().Msg(fmt.Sprint(data...))
}

func (l loggerHandler) Fatal(data ...interface{}) {
	log.Fatal().Msg(fmt.Sprint(data...))
}

func InitLoggerHandler() {
	webgo.LOGHANDLER = &loggerHandler{}
}
