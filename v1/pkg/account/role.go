package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"io"
	"log"
	rolePt "micro-pro/v1/proto/role"
	"reflect"
)

var (
	RoleCli rolePt.RoleClient
)

func (n *node) RoleIndex(req *restful.Request, rsp *restful.Response) {

	nodePId := req.QueryParameter("parentId")
	nodeguid := req.QueryParameter("guid")
	nodeType := req.QueryParameter("type")
	if nodePId == "" {
		log.Print("node index")
		nodePId = "#"
		log.Print(reflect.TypeOf(accPt.NodeReq{}).String())
		response, err := Acl.NodeIndex(context.TODO(), &accPt.NodeReq{
			ParentId: nodePId,
			Guid:     nodeguid,
			Type:     nodeType,
		})
		if err != nil {
			rsp.WriteError(500, err)
		}

		rsp.WriteEntity(response)

	} else {
		log.Print("node children")
		var savedFeatures []*accPt.NodeRsp
		var nilFeatures [0]accPt.NodeRsp
		stream, err := Acl.NodeChildren(context.TODO(), &accPt.NodeReq{
			ParentId: nodePId,
			Guid:     nodeguid,
			Type:     nodeType,
		})
		if err != nil {
			rsp.WriteError(500, err)
		}

		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				log.Println(111)
				break
			}
			if err != nil {
				log.Fatalf("%v.ListFeatures(_) = _, %v", Acl, err)
			}
			log.Println(feature)
			savedFeatures = append(savedFeatures, feature)
		}
		log.Print(len(savedFeatures))
		if len(savedFeatures) == 0 {
			rsp.WriteEntity(nilFeatures)
		} else {
			rsp.WriteEntity(savedFeatures)
		}
	}

}

// /v1/user/accounts/3?type=organization
func (n *node) RoleRead(req *restful.Request, rsp *restful.Response) {
	nodePId := req.QueryParameter("id")
	nodeType := req.QueryParameter("type")
	nodeguid := req.PathParameter("id")
	if nodePId != "" {
		log.Print("node index")
		nodePId = "#"
		log.Print(reflect.TypeOf(accPt.NodeReq{}).String())
		response, err := Acl.NodeRead(context.TODO(), &accPt.NodeReq{
			Guid: nodeguid,
			Type: nodeType,
		})
		if err != nil {
			rsp.WriteError(500, err)
		}

		rsp.WriteEntity(response)
	} else {
		rsp.WriteError(500, "err")
	}
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
