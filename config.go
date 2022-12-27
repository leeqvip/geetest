package geetestbot

type ConfigApi struct {
	AppId   string `mapstructure:"app_id"`
	AppKey  string `mapstructure:"app_key"`
	Timeout int64  `mapstructure:"timeout"`
}

type Config struct {
	Default string               `mapstructure:"default"`
	Apis    map[string]ConfigApi `mapstructure:"apis"`
}
