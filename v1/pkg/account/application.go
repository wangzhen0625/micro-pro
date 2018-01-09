package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"io"
	"log"
	appPt "micro-pro/v1/proto/application"
	"reflect"
)

var (
	AppCli appPt.AppClient
)

type LiAttr struct {
	Guid string `json:"guid" default:"0"`
	Type string `json:"type" description:"id" default:"0"`
}
type User struct {
	Children bool   `json:"children" description:"children"  default:"0"`
	Id       string `json:"id" description:"id" default:"0"`
	Liattr   LiAttr `json:"li_attr" default:"0"`
	Parent   string `json:"parent" default:"0"`
	Text     string `json:"text" default:"0"`
	Type     string `json:"type" default:"0"`
}

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
func (n *node) NodeIndex(req *restful.Request, rsp *restful.Response) {

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
func (n *node) NodeRead(req *restful.Request, rsp *restful.Response) {
	nodeType := req.QueryParameter("type")
	nodeguid := req.PathParameter("id")
	if nodePId !== "" {
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
	}else{
		rsp.WriteError(500, "err")
	}
}
func (n *node) NodeSave(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
func (n *node) NodePatch(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
func (n *node) NodeDelete(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{})
}
