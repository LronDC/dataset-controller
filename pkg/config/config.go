package config

import "github.com/spf13/viper"

var config *viper.Viper

func init() {
	config = viper.New()
	config.AutomaticEnv()
	config.BindEnv("level", "LOG_LEVEL")
	config.SetDefault("level", "debug")
	// bind COMPLETE_NOTIFY_URL env var
	config.BindEnv("complete_notify_url", "COMPLETE_NOTIFY_URL")

}

func GetLevel() string {
	return config.GetString("level")
}

// GetCompleteNotifyURL fetch COMPLETE_NOTIFY_URL env var
func GetCompleteNotifyURL() string {
	return config.GetString("complete_notify_url")
}
