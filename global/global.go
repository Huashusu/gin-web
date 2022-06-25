package global

import (
	"gin-web/config"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

var (
	VIPER  *viper.Viper
	CONFIG *config.Server
	LOG    *zap.Logger
)
