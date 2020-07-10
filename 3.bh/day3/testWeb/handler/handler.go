package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	testWeb "testWeb/proto/testSrv"   // 最前面的 testWeb 为别名
)

func TestWebCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	/*
	修改 默认生成的 NewTestWebService 函数名为:
	实际的函数名: NewTestSrvService
	*/
	// call the backend service
	testWebClient := testWeb.NewTestSrvService("go.micro.srv.testSrv", client.DefaultClient)
	rsp, err := testWebClient.Call(context.TODO(), &testWeb.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
