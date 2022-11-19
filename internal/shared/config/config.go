package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"NAME"`
	SSLMode  string `mapstructure:"SSLMODE"`
	Timezone string `mapstructure:"TIMEZONE"`
}

type ServicesConfig struct {
	TourneyMakerURL    string `mapstructure:"TOURNEY_MAKER_URL"`
	TourneyRegistryURL string `mapstructure:"TOURNEY_REGISTRY_URL"`
}
type Config struct {
	Server   ServerConfig   `mapstructure:"SERVER"`
	Database DatabaseConfig `mapstructure:"DATABASE"`
	Services ServicesConfig `mapstructure:"SERVICES"`
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
