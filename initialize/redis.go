package initialize

import (
	"github.com/go-redis/redis"
	"meteor/global"
	"meteor/httptrack/syslog"
	"runtime"
)

// redis 配置
func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisCfg.MasterName,
		SentinelAddrs: redisCfg.SentinelAddrs, // no password set
		DB:            redisCfg.DB,            // use default DB
		PoolSize:      redisCfg.PoolSize * runtime.NumCPU(),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		syslog.GetLogger().Error("redis connect ping failed, err:", err)
	} else {
		syslog.GetLogger().Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
