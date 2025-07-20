// config/env.go
package config

import "github.com/spf13/viper"

type Config struct {
    ServerPort string
}

func Load() *Config {
    viper.SetDefault("PORT", "50051")
    viper.BindEnv("PORT")
    return &Config{ServerPort: viper.GetString("PORT")}
}
