package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"log"
	accPt "micro-pro/v1/proto/account"
)

var (
	acl accPt.AccountClient
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
  {children: true, id: "3", li_attr:{guid:"3",type: "organization"}, parent: "#", text: "topsec",type: "organization"},
  {children: true, id: "3", li_attr:{guid:"3",type: "organization"}, parent: "#", text: "topsec",type: "organization"}

  ]
*/
func (n *node) NodeIndex(req *restful.Request, rsp *restful.Response) {
	log.Print("node index")

	nodePId := req.QueryParameter("parentId")
	nodeguid := req.QueryParameter("guid")
	nodeType := req.QueryParameter("type")
	if nodePId == "" {
		nodePId = "#"
	}
	response, err := acl.NodeIndex(context.TODO(), &accPt.NodeReq{
		ParentId: nodePId,
		Guid:     nodeguid,
		Type:     nodeType,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}
func (n *node) NodeRead(req *restful.Request, rsp *restful.Response) {
	liattr := LiAttr{Guid: "3", Type: "organization"}
	log.Print("Received Rest.Test API request")
	rsp.WriteEntity(User{Children: true, Id: "3", Liattr: liattr, Parent: "#", Text: "topsec", Type: "organization"})
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
