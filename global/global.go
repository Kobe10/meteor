package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"meteor/config"
)

var (
	GVA_DB    *gorm.DB
	GVA_REDIS *redis.Client
	// 总配置文件
	GVA_VP *viper.Viper
	// 各个环境配置文件
	GVA_ENV_VP *viper.Viper
	GVA_CONFIG config.Server
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)
