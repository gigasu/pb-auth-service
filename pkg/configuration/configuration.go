package configuration

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Grpc       Grpc       `yaml:"grpc"`
	Profile    string     `yaml:"profile"`
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}
type Grpc struct {
	AuthSvcUrl string `yaml:"authSvcUrl"`
}

type Datasource struct {
	DbType   string `yaml:"dbType"`
	Url      string `yaml:"url"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}
type Server struct {
	Port string `yaml:"port"`
}

func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}
func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("auth GO_PROFILE: " + profile)
	return profile
}
func setRuntimeConfig(profile string) {
	var err error
	viper.AddConfigPath(".")
	viper.AddConfigPath("./pkg/configuration")
	viper.AddConfigPath("./pb-api-gateway/pkg/configuration")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&RuntimeConf)
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	viper.WatchConfig()
}
