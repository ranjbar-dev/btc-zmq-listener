package config

import (
	"strings"

	"github.com/spf13/viper"
)

var data map[string]interface{}

func init() {

	v := viper.New()

	v.SetConfigType("yaml")

	v.AddConfigPath("config")

	v.AutomaticEnv()

	v.SetEnvPrefix("engine")

	v.SetEnvKeyReplacer(strings.NewReplacer(":", "_"))

	err := v.ReadInConfig()
	if err != nil {

		panic("error in reading config file: " + err.Error())
	}

	data = v.AllSettings()
}

func Timezone() string {

	return data["app"].(map[string]interface{})["timezone"].(string)
}

func GatewayHost() string {

	return data["gateway"].(map[string]interface{})["host"].(string)
}

func GatewayPort() string {

	return data["gateway"].(map[string]interface{})["port"].(string)
}

func GatewayWhiteListIps() []string {

	return data["gateway"].(map[string]interface{})["whitelist_ip"].([]string)
}

func ZmqAddress() string {

	return data["zmq"].(map[string]interface{})["address"].(string)
}

func TelegramBotToken() string {

	return data["telegram"].(map[string]interface{})["token"].(string)
}

func TelegramChatID() int64 {

	return data["telegram"].(map[string]interface{})["chat_id"].(int64)
}
