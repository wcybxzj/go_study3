package main

import (
	"github.com/micro/go-micro/util/log"
	"net/http"

	"github.com/micro/go-micro/web"
	"testWeb/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.testWeb"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/testWeb/call", handler.TestWebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
