package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"log"
	rolePt "micro-pro/v1/proto/role"
	"reflect"
)

var (
	RoleCli rolePt.RoleClient
)

func (n *node) RoleIndex(req *restful.Request, rsp *restful.Response) {

	log.Print("node index")
	log.Print(reflect.TypeOf(rolePt.RoleReq{}).String())
	response, err := RoleCli.RoleIndex(context.TODO(), &rolePt.RoleReq{
		Id:   1,
		Name: "2",
	})
	if err != nil {
		rsp.WriteError(500, err)
	}
	rsp.WriteEntity(response)

}

// /v1/user/accounts/3?type=organization
func (n *node) RoleRead(req *restful.Request, rsp *restful.Response) {
	log.Print("node index")
	log.Print(reflect.TypeOf(rolePt.RoleReq{}).String())
	response, err := RoleCli.RoleRead(context.TODO(), &rolePt.RoleReq{
		Id:   1,
		Name: "2",
	})
	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}
func (n *node) RoleSave(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
}
func (n *node) RolePatch(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
}
func (n *node) RoleDelete(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
}
