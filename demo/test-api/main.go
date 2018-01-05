package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
	"golang.org/x/net/context"
	hello "micro-pro/v1/test/test-srv/hello"
)

type Say struct {
	Client hello.SayClient
}

func (s *Say) SayHello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank")
	}

	response, err := s.Client.SayHello(ctx, &hello.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.v1.test"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewSayClient("go.micro.srv.v1.test", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
