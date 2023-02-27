package config

// 敏感词配置
type Eunomia struct {
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	Domain      string `mapstructure:"domain" json:"domain" yaml:"domain"`
	AccessToken string `mapstructure:"access_token" json:"accessToken" yaml:"access_token"`
	Sensitive   string `mapstructure:"sensitive" json:"sensitive" yaml:"sensitive"`
	Oauth       string `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
}
