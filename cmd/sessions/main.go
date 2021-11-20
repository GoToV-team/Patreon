package main

import (
	"flag"
	"os"
	"patreon/internal/app"
	sessionServer "patreon/internal/microservices/auth/delivery/grpc/server"
	"patreon/internal/microservices/auth/sessions/repository"
	sessions_manager2 "patreon/internal/microservices/auth/sessions/sessions_manager"
	"patreon/pkg/utils"

	grpc2 "google.golang.org/grpc"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	config := app.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}
	logger, CloseLogger := utils.NewLogger(config, true, "session_microservice")
	defer CloseLogger()
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		os.Exit(1)
	}
	logger.SetLevel(level)

	sessionRedisPool := utils.NewRedisPool(config.ServerRepository.SessionRedisUrl)
	logger.Info("Session-service new redis pool create")

	conn, err := sessionRedisPool.Dial()
	if err != nil {
		logger.Fatal(err)
	}
	if err = conn.Close(); err != nil {
		logger.Fatal(err)
	}
	logger.Info("Session-service new redis pool success check")

	grpc := grpc2.NewServer()
	sessionRepository := repository.NewRedisRepository(sessionRedisPool, logger)
	logger.Info("Session-service create repository")

	server := sessionServer.NewAuthGRPCServer(logger, grpc, sessions_manager2.NewSessionManager(sessionRepository))
	if err = server.StartGRPCServer(config.Microservices.SessionServerUrl); err != nil {
		logger.Fatalln(err)
	}
	logger.Info("Session-service was stopped")

}
