package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/emicklei/go-restful"
	"github.com/liudng/godump"
	// "github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	// "golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

type Rest struct{}

// User is just a sample type
type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}

func (r *Rest) Test(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
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

func (r *Rest) Test2(req *restful.Request, rsp *restful.Response) {
	godump.Dump(123)
	usr := User{}
	err := req.ReadEntity(&usr)
	godump.Dump(usr)
	// here you would create the user with some persistence system
	if err == nil {
		rsp.StatusCode()
		rsp.WriteEntity(usr)
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}
}

func (r *Rest) Test3(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test3 API request")

	age := req.PathParameter("age")
	rsp.WriteEntity(age)
}
func (r Rest) NewContainer() *restful.WebService {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/v1/rest")
	ws.Route(ws.GET("/").To(r.Test))
	ws.Route(ws.POST("/{name}").To(r.Test2))
	ws.Route(ws.GET("/{name}/test3/{age}").To(r.Test3))
	return ws
}
func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.v1.rest"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)

	service.Init()

	rest := new(Rest)
	// Register Handler
	wc := restful.NewContainer()
	wc.Add(rest.NewContainer())
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
