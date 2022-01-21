package main

import (
	"fmt"
	"github.com/spf13/viper"
	log "walletService/lib/logging"
	"walletService/server/amqprpc/amqprpc"
)

func main() {
	host:= viper.GetString("mq.host")
	port := viper.GetInt("mq.port")
	user:= viper.GetString("mq.user")
	pwd := viper.GetString("mq.password")
	vhost := viper.GetString("mq.vhost")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",user, pwd, host, port, vhost)
	exchange := viper.GetString("mq.exchange")
	prefetchCount := viper.GetInt("mq.prefetchCount")
	reconnectInterval := viper.GetInt("mq.reconnectInterval")

	client, err := amqprpc.NewClient(&amqprpc.Config{
		Dsn:               dsn,
		ClientTimeout:     10,
		ReconnectInterval: reconnectInterval,
		Log:               log.StandardLog(),
		Exchange:          exchange,
		PrefetchCount:     prefetchCount,
		Serializer: new(amqprpc.JsonSerializer),
	})

	if err != nil {
		log.Error(err)
	}

	for i := 0; i < 100; i++ {
		params := amqprpc.Args{A: 5, B: i}

		var result amqprpc.Result
		if err := client.Call(amqprpc.Six007rpcApi, params, &result); err != nil {
			log.Error(err)
		}
		log.Error(result)
	}

	if err := client.Close(); err != nil {
		log.Error(err)
	}
}
