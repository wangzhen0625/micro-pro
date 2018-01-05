package main

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-web"
	"log"
	"micro-pro/v1/pkg/account"
	"time"
)

type Account struct{}

func (r Account) NewContainer() *restful.WebService {
	node := account.CreateNode()
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/v1/account")

	ws.Route(ws.GET("/nodes").To(node.NodeIndex))          //index
	ws.Route(ws.GET("/nodes/{id}").To(node.NodeRead))      //read
	ws.Route(ws.POST("/nodes").To(node.NodeSave))          //save
	ws.Route(ws.PATCH("/nodes/{id}").To(node.NodePatch))   //patch
	ws.Route(ws.DELETE("/nodes/{id}").To(node.NodeDelete)) //delete

	return ws
}
func main() {

	// rest := account.CreateRest()
	// Register Handler
	account := new(Account)
	wc := restful.NewContainer()
	wc.Add(account.NewContainer())

	// Create service
	service := web.NewService(
		web.Name("go.micro.api.v1.account"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)
	service.Init()
	service.Handle("/", wc)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
