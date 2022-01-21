package amqprpc

import (
	"encoding/json"
	"walletService/lib/logging"
)

type JsonSerializer struct {}
func (s *JsonSerializer)Marshal(v interface{})([]byte, error) {
	return json.Marshal(v)
}
func (s *JsonSerializer)Unmarshal(data []byte, v interface{})  error{
	return json.Unmarshal(data, v)
}
func (s *JsonSerializer)GetContentType() string {
	return "application/json"
}

type JsonMethod struct {
	serializer  Serializer
}

func (m *JsonMethod) GetName() string {
	return "json"
}

func (m *JsonMethod) Setup(serializer Serializer) error {
	m.serializer = serializer
	return nil
}

func (m *JsonMethod) Cleanup() error {
	return nil
}

func (m *JsonMethod) Call(body []byte) (interface{}, *RPCError) {
	var params Args
	if err := m.serializer.Unmarshal(body, &params); err != nil {
		return nil, &RPCError{
			Err: ErrorData{
				Type:    "UnmarshalError",
				Message: err.Error(),
			},
		}
	}
	res := params.A * params.B
	logging.Printf("Result: %d", res)
	return &Result{Result: res}, nil
}
