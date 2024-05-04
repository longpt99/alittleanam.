package configs

import (
	"ala-coffee-notification/models"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Env models.Env

func InitConfig() error {
	pwd, _ := os.Getwd()

	viper.SetConfigFile(fmt.Sprintf("%s/configs/env/.env", pwd))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Env)

	return err
}
