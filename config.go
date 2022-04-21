package service

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var Config = &Configuration{
	Host:      "",
	Port:      "8080",
	LogEncode: "console",
	LogLevel:  "debug",
	Database: Database{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "123456",
		Name:     "pda",
	},
}

type Database struct {
	Addr         string `json:"addr" yaml:"addr"`
	User         string `json:"user" yaml:"user"`
	Password     string `json:"password" yaml:"password"`
	Name         string `json:"name" yaml:"name"`
	StmtLog      bool   `json:"stmt_log" yaml:"stmt_log"`
	DialTimeout  uint   `json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  uint   `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout uint   `json:"write_timeout" yaml:"write_timeout"`
	PoolSize     uint   `json:"pool_size" yaml:"pool_size"`
}

type Configuration struct {
	Host            string   `json:"host" yaml:"host"`
	Port            string   `json:"port" yaml:"port"`
	SessionHashKey  string   `json:"session_hash_key" yaml:"session_hash_key"`
	SessionBlockKey string   `json:"session_block_key" yaml:"session_block_key"`
	LogLevel        string   `json:"log_level" yaml:"log_level"`
	LogFile         string   `json:"log_file" yaml:"log_file"`
	LogEncode       string   `json:"log_encode" yaml:"log_encode"`
	Database        Database `json:"database" yaml:"database"`
}

const configFileName = "config.yml"

func LoadDefaultYAMLConfig() error {
	location, err := os.Executable()
	if err != nil {
		return err
	}
	name := filepath.Join(filepath.Dir(location), configFileName)
	if fileExist(name) {
		return parseFile(name)
	}
	location, err = os.Getwd()
	if err != nil {
		return err
	}
	name = filepath.Join(location, configFileName)
	if fileExist(name) {
		return parseFile(name)
	}
	return fmt.Errorf("file [%s] not exist", name)
}

func parseFile(name string) error {
	file, err := os.Open(filepath.Clean(name))
	if err != nil {
		return err
	}
	return yaml.NewDecoder(file).Decode(&Config)
}

func fileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
