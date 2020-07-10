package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	testSrv "testSrv/proto/testSrv"
)

type TestSrv struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *TestSrv) Call(ctx context.Context, req *testSrv.Request, rsp *testSrv.Response) error {
	log.Log("Received TestSrv.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
