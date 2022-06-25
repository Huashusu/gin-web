package main

import (
	"gin-web/core"
	"gin-web/global"

	"go.uber.org/zap"
)

func main() {

	global.VIPER = core.Viper()

	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)

	core.RunServer()
}
