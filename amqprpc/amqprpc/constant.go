package amqprpc

const (
	SeamExchange = "spgs_seamless_exchange"
	Six007rpcApi = "6007_rpc_api"
)

type Args struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result struct {
	Result int `json:"result"`
}





