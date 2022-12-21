package main

import (
	"task_rest/middleware"
	"task_rest/model"
	"task_rest/net"
)

func main() {
	middleware.Logs.Info().Msgf("start main")
	middleware.LockDebug(model.ConfigFile.Api.DebugMode)

	middleware.Logs.Info().Msgf("start net client")
	net.StartNet()

	middleware.Logs.Info().Msgf("run net client")
	net.RunNet()

	middleware.Logs.Info().Msgf("close app")
}
