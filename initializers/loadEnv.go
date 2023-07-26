package initializers

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

// Allows for different configurations
func LoadConfig(path string) (config Config, err error) {

	// Viper handles configuration of applicaton
	// Takes in the app.env
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// We can give it different app.env
	viper.AddConfigPath(path)

	// Checks if there are any files that it can automatically use
	// in the project
	viper.AutomaticEnv()

	// Reads the config into viper
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshals config, returns address of configuration
	err = viper.Unmarshal(&config)
	return
}
