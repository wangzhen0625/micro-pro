package account

import (
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
	"io"
	"log"
	nodePt "micro-pro/v1/proto/node"
	"net/http"
	"reflect"
)

var (
	NodeCli nodePt.NodeClient
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
		log.Print(reflect.TypeOf(nodePt.NodeReq{}).String())
		response, err := NodeCli.NodeIndex(context.TODO(), &nodePt.NodeReq{
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
		var savedFeatures []*nodePt.NodeRsp
		var nilFeatures [0]nodePt.NodeRsp
		stream, err := NodeCli.NodeChildren(context.TODO(), &nodePt.NodeReq{
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
				log.Fatalf("%v.ListFeatures(_) = _, %v", NodeCli, err)
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

// /v1/user/accounts/3?type=organization  类型用来查找属性
func (n *node) NodeRead(req *restful.Request, rsp *restful.Response) {
	log.Print("NodeRead")
	nodeguid := req.PathParameter("id")
	nodeType := req.QueryParameter("type")
	log.Print(nodeguid)
	log.Print(nodeType)
	response, err := NodeCli.NodeRead(context.TODO(), &nodePt.NodeReq{
		Guid: nodeguid,
		Type: nodeType,
	})
	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

///v1/user/accounts post
func (n *node) NodeSave(req *restful.Request, rsp *restful.Response) {
	log.Print("NodeSave")

	node := nodePt.NodeInfo{}
	err := req.ReadEntity(&node)
	if err == nil {
		rsp.WriteHeaderAndEntity(http.StatusCreated, node)
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}

	response, err := NodeCli.NodeSave(context.TODO(), &node)
	if err == nil {
		rsp.WriteHeaderAndEntity(http.StatusOK, response)
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}
}

// /v1/user/accounts/id patch
func (n *node) NodePatch(req *restful.Request, rsp *restful.Response) {
	log.Print("NodePatch")
	nodeguid := req.PathParameter("id")
	node := nodePt.NodeInfo{}
	node.Id = nodeguid
	err := req.ReadEntity(&node)
	if err == nil {
		rsp.WriteHeaderAndEntity(http.StatusCreated, node)
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}

	response, err := NodeCli.NodePatch(context.TODO(), &node)
	if err == nil {
		rsp.WriteHeaderAndEntity(http.StatusOK, response)
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}
}

// /v1/user/accounts/id delete 删除返回空数据
func (n *node) NodeDelete(req *restful.Request, rsp *restful.Response) {
	log.Print("NodeDelete api")
	nodeguid := req.PathParameter("id")
	_, err := NodeCli.NodeDelete(context.TODO(), &nodePt.NodeReq{
		Guid: nodeguid,
	})
	if err == nil {
		rsp.WriteHeaderAndEntity(http.StatusNoContent, "")
	} else {
		rsp.WriteError(http.StatusInternalServerError, err)
	}
}
