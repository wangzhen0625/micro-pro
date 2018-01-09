package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"log"
	appPt "micro-pro/v1/proto/application"
	"reflect"
)

var (
	AppCli appPt.AppClient
)

/*
index
根节点 /v1/account/nodes?parentId=&guid=&type=
  {children: true, id: "3", li_attr:{guid:"3",type: "organization"}, parent: "#", text: "topsec",type: "organization"}
其他结点 /v1/account/nodes?parentId=3&guid=3&type=organization
  [
  {children: true, id: "4", li_attr:{guid:"3",type: "organization"}, parent: "3", text: "topsec",type: "organization"},
  {children: true, id: "5", li_attr:{guid:"3",type: "organization"}, parent: "3", text: "topsec",type: "organization"}

  ]
*/
func (n *node) AppIndex(req *restful.Request, rsp *restful.Response) {

	nodeguid := req.QueryParameter("guid")
	log.Print("node index")
	log.Print(reflect.TypeOf(appPt.AppReq{}).String())
	response, err := AppCli.AppIndex(context.TODO(), &appPt.AppReq{
		Id:   1,
		Name: nodeguid,
	})
	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)

}

// /v1/user/accounts/3?type=organization
func (n *node) AppRead(req *restful.Request, rsp *restful.Response) {
	nodeguid := req.PathParameter("id")
	if nodeguid != "" {
		log.Print("node index")
		log.Print(reflect.TypeOf(appPt.AppReq{}).String())
		response, err := AppCli.AppRead(context.TODO(), &appPt.AppReq{
			Id: 1,
		})
		if err != nil {
			rsp.WriteError(500, err)
		}

		rsp.WriteEntity(response)
	}
}
func (n *node) AppSave(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
func (n *node) AppPatch(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
func (n *node) AppDelete(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
