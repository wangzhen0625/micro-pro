package main

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"log"
	"micro-pro/v1/pkg/account"
	// appPt "micro-pro/v1/proto/application"
	nodePt "micro-pro/v1/proto/node"
	// rolePt "micro-pro/v1/proto/role"
	"time"
)

type Account struct{}

func (r Account) NewContainer() *restful.WebService {

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/v1/account")

	account.NodeCli = nodePt.NewNodeClient("go.micro.srv.v1.node", client.DefaultClient)
	node := account.CreateNode()
	ws.Route(ws.GET("/nodes").To(node.NodeIndex))          //index
	ws.Route(ws.GET("/nodes/{id}").To(node.NodeRead))      //read
	ws.Route(ws.POST("/nodes").To(node.NodeSave))          //save
	ws.Route(ws.POST("/nodes/signup").To(node.NodeSignUp)) //signup
	ws.Route(ws.PUT("/nodes/{id}").To(node.NodePatch))     //patch
	ws.Route(ws.DELETE("/nodes/{id}").To(node.NodeDelete)) //delete

	/*	account.RoleCli = rolePt.NewRoleClient("go.micro.srv.v1.role", client.DefaultClient)
		role := account.CreateRole()
		ws.Route(ws.GET("/roles").To(role.RoleIndex))          //index
		ws.Route(ws.GET("/roles/{id}").To(role.RoleRead))      //read
		ws.Route(ws.POST("/roles").To(role.RoleSave))          //save
		ws.Route(ws.PATCH("/roles/{id}").To(role.RolePatch))   //patch
		ws.Route(ws.DELETE("/roles/{id}").To(role.RoleDelete)) //delete

		account.AppCli = appPt.NewAppClient("go.micro.srv.v1.app", client.DefaultClient)
		app := account.CreateApp()
		ws.Route(ws.GET("/applications").To(app.AppIndex))          //index
		ws.Route(ws.GET("/applications/{id}").To(app.AppRead))      //read
		ws.Route(ws.POST("/applications").To(app.AppSave))          //save
		ws.Route(ws.PATCH("/applications/{id}").To(app.AppPatch))   //patch
		ws.Route(ws.DELETE("/applications/{id}").To(app.AppDelete)) //delete*/

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
