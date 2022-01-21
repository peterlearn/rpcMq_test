package amqprpc

import (
	"fmt"
	"github.com/spf13/viper"
	log "walletApi/lib/logging"
	"walletApi/server/amqprpc/amqprpc"
)

var (
	server     = new(amqprpc.Server)
	RpcQueList = make([]string, 0)
	err        error
)

func init() {
	queList := viper.GetStringSlice("rpc_que")
	for _, v := range queList {
		RpcQueList = append(RpcQueList, fmt.Sprintf("%v_rpc_api", v))
	}
}

func Start() {
	host := viper.GetString("mq.host")
	port := viper.GetInt("mq.port")
	user := viper.GetString("mq.user")
	pwd := viper.GetString("mq.password")
	vhost := viper.GetString("mq.vhost")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", user, pwd, host, port, vhost)
	exchange := viper.GetString("mq.exchange")
	prefetchCount := viper.GetInt("mq.prefetchCount")
	reconnectInterval := viper.GetInt("mq.reconnectInterval")
	server, err = amqprpc.NewServer(&amqprpc.Config{
		Dsn: dsn,
		Exchange:          exchange,
		AutoDelete:        false,
		IsDurable:         true,
		ReconnectInterval: reconnectInterval,
		Log:               log.StandardLog(),
		PrefetchCount:     prefetchCount,
	})

	if err != nil {
		log.Error(err)
	}

	meth := new(amqprpc.JsonMethod)
	registry := amqprpc.NewRegistry("game_rpc_api")

	// 启动所有的rpc队列
	for _, que := range RpcQueList {
		registry.AddMethod(que, meth)
	}

	//registry.AddMethod(meth)
	server.AddRegistry(registry)

	if err := server.Setup(); err != nil {
		log.Error(err)
	}
}

func Stop() {
	if err := server.Close(); err != nil {
		log.Error(err)
	}
}

func StartConsumer(que string, num int) error {
	err := server.StartComsumer(que, num)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
