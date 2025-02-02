package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var data *viper.Viper

func init() {

	data = viper.New()

	data.SetConfigType("yaml")

	if len(os.Args) > 1 {

		data.SetConfigFile(os.Args[1])
	} else {

		data.SetConfigName("config")
	}

	data.AutomaticEnv()

	data.SetEnvPrefix("btczmq")

	data.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := data.ReadInConfig()
	if err != nil {

		panic("error in reading config file: " + err.Error())
	}

	data.AllSettings()
}

func Timezone() string {

	return data.GetString("app.timezone")
}

func GatewayHost() string {

	return data.GetString("gateway.host")
}

func GatewayPort() string {

	return data.GetString("gateway.port")
}

func GatewayWhiteListIps() []string {

	return data.GetStringSlice("gateway.whitelist_ip")
}

func ZmqAddress() string {

	return data.GetString("zmq.address")
}

func TelegramBotToken() string {

	return data.GetString("telegram.token")
}

func TelegramChatID() int64 {

	return data.GetInt64("telegram.chat_id")
}

func BitcoinProtocol() string {

	return data.GetString("bitcoin.protocol")
}

func BitcoinHost() string {

	return data.GetString("bitcoin.host")
}

func BitcoinPort() string {

	return data.GetString("bitcoin.port")
}

func BitcoinUser() string {

	return data.GetString("bitcoin.user")
}

func BitcoinPass() string {

	return data.GetString("bitcoin.password")
}
