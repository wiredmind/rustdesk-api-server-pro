package config

import (
	"fmt"
	"os"
	"path"
	"rustdesk-api-server-pro/util"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	DebugMode  bool        `yaml:"debugMode"`
	Db         *DbConfig   `yaml:"db"`
	SignKey    string      `yaml:"signKey"`
	HttpConfig *HttpConfig `yaml:"httpConfig"`
	SmtpConfig *SmtpConfig `yaml:"smtpConfig"`
	JobsConfig *JobsConfig `yaml:"jobsConfig"`
}

type DbConfig struct {
	Driver   string `yaml:"driver"`
	Dsn      string `yaml:"dsn"`
	TimeZone string `yaml:"timeZone"`
	ShowSql  bool   `yaml:"showSql"`
}

type HttpConfig struct {
	PrintRequestLog bool   `yaml:"printRequestLog"`
	Port            string `yaml:"port"`
	StaticDir       string `yaml:"staticdir"`
}

type SmtpConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Encryption string `yaml:"encryption"` // none ssl/tls starttls
	From       string `yaml:"from"`
}

type DeviceCheckJob struct {
	Duration int `yaml:"duration"`
}

type JobsConfig struct {
	DeviceCheckJob *DeviceCheckJob `yaml:"deviceCheckJob"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		DebugMode: false,
		Db: &DbConfig{
			Driver:   "sqlite",
			Dsn:      "./server.db",
			ShowSql:  true,
			TimeZone: "Asia/Shanghai",
		},
		HttpConfig: &HttpConfig{
			Port:      ":8080",
			StaticDir: "dist",
		},
		SignKey: util.RandomString(32),
		JobsConfig: &JobsConfig{
			DeviceCheckJob: &DeviceCheckJob{
				Duration: 30,
			},
		},
	}
}

func GetServerConfig() *ServerConfig {
	cfg := GetDefaultServerConfig()
	bytes, err := os.ReadFile(yamlFile)
	if os.IsNotExist(err) {
		WriteServerConfig(cfg)
		return cfg
	}
	if err != nil {
		panic(fmt.Errorf("read server config %q: %w", yamlFile, err))
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		panic(fmt.Errorf("parse server config %q: %w", yamlFile, err))
	}
	return cfg
}

func WriteServerConfig(cfg *ServerConfig) {
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		panic(fmt.Errorf("marshal server config: %w", err))
	}
	if err := os.WriteFile(yamlFile, bytes, 0600); err != nil {
		panic(fmt.Errorf("write server config %q: %w", yamlFile, err))
	}
}
