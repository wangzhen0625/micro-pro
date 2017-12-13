package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/emicklei/go-restful"
	// "github.com/liudng/godump"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-web"
	"golang.org/x/net/context"
	"log"
	hello "micro-pro/v1/test/test-srv/hello"
)

type Say struct{}

// User is just a sample type
type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

var (
	cl hello.SayClient
)

func (s *Say) Anything(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Anything API request")
	/*rsp.WriteEntity(map[string]string{
		"message": "Hi, this is the Greeter API",
	})*/
	// b, err := json.Marshal(User{ID: "111", Name: "wzzz", Age: 12})
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// godump.Dump(User{ID: "111", Name: "wzzz", Age: 12})
	rsp.WriteEntity(User{ID: "111", Name: "wzzz", Age: 12})
}

func (s *Say) Hello(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Hello API request")

	name := req.PathParameter("name")

	response, err := cl.SayHello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}
func (s Say) NewContainer() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/v1/test")
	ws.Route(ws.GET("/").To(s.Anything))
	ws.Route(ws.GET("/{name}").To(s.Hello))
	return ws
}
func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.v1.test"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = hello.NewSayClient("go.micro.srv.v1.test", client.DefaultClient)

	// Create RESTful handler
	say := new(Say)
	// Register Handler
	wc := restful.NewContainer()
	wc.Add(say.NewContainer())
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
