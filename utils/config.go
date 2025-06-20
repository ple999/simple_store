package utils

import(
	viper "github.com/spf13/viper"
	"time"
)

type Config struct{
	DBConnection string `mapstructure:"DB_CONNECTION"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	PasetoSymetricKey string `mapstructure:"PASETO_SYMETRIC_KEY"`
	TokenExpiredTime time.Duration `mapstructure:"TOKEN_EXPIRED_TIME"`
}

func LoadConfig(path string) (config Config,err error){
	viper.AddConfigPath(path);
	viper.SetConfigName("app");
	viper.SetConfigType("env");

	viper.AutomaticEnv();
	err = viper.ReadInConfig();
	if err!=nil{
		return;
	}

	err = viper.Unmarshal(&config);
	return

}