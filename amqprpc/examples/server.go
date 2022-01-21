package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	_ "walletService/constant"
	log "walletService/lib/logging"
	"walletService/server/amqprpc/amqprpc"
)


//func init() {
//	log.SetOutput(os.Stdout)
//	log.SetLevel(log.DebugLevel)
//}

var (
	RpcQueList = make([]string, 0)
)



func main() {
	queList := viper.GetStringSlice("rpc_que")
	for _, v := range queList {
		RpcQueList = append(RpcQueList, fmt.Sprintf("%v_rpc_api", v))
	}
	host:= viper.GetString("mq.host")
	port := viper.GetInt("mq.port")
	user:= viper.GetString("mq.user")
	pwd := viper.GetString("mq.password")
	vhost := viper.GetString("mq.vhost")
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",user, pwd, host, port, vhost)
	exchange := viper.GetString("mq.exchange")
	prefetchCount := viper.GetInt("prefetchCount")
	reconnectInterval := viper.GetInt("reconnectInterval")

	server, err := amqprpc.NewServer(&amqprpc.Config{
		//Dsn:               "amqp://admin:admin@localhost:5672/mq-api",
		Dsn: dsn,
		//Exchange:          amqprpc.SeamExchange,
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

	for _, que := range RpcQueList {
		registry.AddMethod(que, meth)
	}

	//registry.AddMethod(meth)
	server.AddRegistry(registry)

	if err := server.Setup(); err != nil {
		log.Error(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGILL, syscall.SIGINT, syscall.SIGQUIT)
	<-ch

	if err := server.Close(); err != nil {
		log.Error(err)
	}
}
