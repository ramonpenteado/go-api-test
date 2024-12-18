package app

import (
	"test/api/internal/app/config/environment"
	server "test/api/internal/app/http/server"
)

func Init() {
	environment.LoadEnv()
	server.Init()
}
