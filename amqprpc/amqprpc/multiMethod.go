package amqprpc

import "walletApi/lib/logging"

type MultiplyMethod struct {
	serializer Serializer
}

func (m *MultiplyMethod) GetName() string {
	return "multiply"
}

func (m *MultiplyMethod) Setup(serializer Serializer) error {
	m.serializer = serializer
	return nil
}

func (m *MultiplyMethod) Cleanup() error {
	return nil
}

func (m *MultiplyMethod) Call(body []byte) (interface{}, *RPCError) {
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

