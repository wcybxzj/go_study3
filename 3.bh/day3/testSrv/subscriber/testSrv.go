package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	testSrv "testSrv/proto/testSrv"
)

type TestSrv struct{}

func (e *TestSrv) Handle(ctx context.Context, msg *testSrv.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *testSrv.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
