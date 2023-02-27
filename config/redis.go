package config

type Redis struct {
	DB            int      `mapstructure:"db" json:"db" yaml:"db"`
	MasterName    string   `mapstructure:"master" json:"master" yaml:"master"`
	PoolSize      int      `mapstructure:"PoolSize" json:"PoolSize" yaml:"PoolSize"`
	SentinelAddrs []string `mapstructure:"sentinel" json:"sentinel" yaml:"sentinel"`
}
