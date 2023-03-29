package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	fmt.Println("read configuration content......")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	//viper.SetConfigFile("./configuration")
	viper.AddConfigPath("./configuration")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read configuration file error," + err.Error())
	}
	port := viper.GetInt("server.port")
	fmt.Println("server.port ", port)

}
