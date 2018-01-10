package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"log"
	nodePt "micro-pro/v1/proto/node"
	"time"
)

type Node struct{}

var (
	mydb *gorm.DB
	err  error
)

type NodeInfo struct {
	Id         string
	LoginId    string
	Type       string
	AuthMethod string
	ParentNode string
	Name       string
	Password   string
	Code       string
	State      string
	Email      string
	Gender     string
	Rank       string
	Birthday   time.Time
	Tel        string
	MobileTel  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	/*CreatedAt  time.Time  `gorm:"column:create_time"`
	UpdatedAt  time.Time  `gorm:"column:update_time"`
	DeletedAt  *time.Time `gorm:"column:delete_time"`*/
}

func (a *Node) NodeIndex(ctx context.Context, req *nodePt.NodeReq, rsp *nodePt.NodeRsp) error {
	log.Print("Received srv account.NodeIndex request")
	var node NodeInfo
	mydb.Where(&NodeInfo{ParentNode: req.ParentId}).Select("id,parent_node,type,login_id").First(&node)
	log.Print(node)

	//  {children: true, id: "3", li_attr:{guid:"3",type: "organization"}, parent: "#", text: "topsec",type: "organization"}
	liattr := nodePt.LiAttr{}
	liattr.Guid = node.Id
	liattr.Type = node.Type
	if node.Type == "person" {
		rsp.Children = false
	} else {
		rsp.Children = true
	}
	rsp.Id = node.Id
	rsp.LiAttr = &liattr
	rsp.Parent = node.ParentNode
	rsp.Text = node.LoginId
	rsp.Type = node.Type
	return nil
}

/*[
  {children: true, id: "4", li_attr:{guid:"4",type: "organization"}, parent: "3", text: "topsec",type: "organization"},
  {children: true, id: "5", li_attr:{guid:"5",type: "organization"}, parent: "3", text: "topsec",type: "organization"}
   [1=>"organization", 2=>"organizationalUnit", 3=>"person"];
  ]*/
func (a *Node) NodeChildren(ctx context.Context, req *nodePt.NodeReq, stream nodePt.Node_NodeChildrenStream) error {
	var nodes []NodeInfo
	mydb.Where(&NodeInfo{ParentNode: req.ParentId}).Select("id,parent_node,type,login_id").Find(&nodes)

	feature := nodePt.NodeRsp{}
	liattr := nodePt.LiAttr{}
	for _, node := range nodes {
		liattr.Guid = node.Id
		liattr.Type = node.Type
		if node.Type == "person" {
			feature.Children = false
		} else {
			feature.Children = true
		}
		feature.Id = node.Id
		feature.LiAttr = &liattr
		feature.Parent = node.ParentNode
		feature.Text = node.LoginId
		feature.Type = node.Type

		if err := stream.Send(&feature); err != nil {
			return err
		}
	}
	return nil
}

func (a *Node) NodeRead(ctx context.Context, req *nodePt.NodeReq, rsp *nodePt.NodeInfo) error {
	log.Println(req.Type)
	if req.Type == "person" {
		mydb.Where(&nodePt.NodeInfo{Id: req.Guid}).Select("id,login_id,parent_node,auth_method,type,name,email,state,gender,rank").First(rsp)
	} else {
		mydb.Where(&nodePt.NodeInfo{Id: req.Guid}).Select("id,login_id,parent_node,auth_method,type").First(rsp)
	}
	return nil
}

func (a *Node) NodeSave(ctx context.Context, req *nodePt.NodeInfo, rsp *nodePt.NodeInfo) error {
	log.Print("NodeSave")
	//找到相同用户名的数据
	var count int
	if err := mydb.Model(&nodePt.NodeInfo{}).Where("login_id = ?", req.LoginId).Count(&count).Error; err != nil {
		return err
	}
	req.Id = uuid.Must(uuid.NewV4()).String()
	log.Print(count)
	if count > 0 {
		return errors.New("Data already exists")
	}

	data := mydb.Create(req)
	if data.Error != nil {
		return data.Error
	}
	*rsp = *req
	return nil
}

func (a *Node) NodeSignUp(ctx context.Context, req *nodePt.NodeRegister, rsp *nodePt.NodeRegister) error {
	log.Print("NodeSave")
	//找到相同用户名的数据
	var count int
	if err := mydb.Model(&nodePt.NodeRegister{}).Where("login_id = ?", req.LoginId).Count(&count).Error; err != nil {
		return err
	}
	req.Id = uuid.Must(uuid.NewV4()).String()
	log.Print(count)
	if count > 0 {
		return errors.New("Data already exists")
	}

	data := mydb.Create(req)
	if data.Error != nil {
		return data.Error
	}
	*rsp = *req
	return nil
}

func (a *Node) NodePatch(ctx context.Context, req *nodePt.NodeInfo, rsp *nodePt.NodeInfo) error {
	log.Print("NodePatch")

	if err := mydb.Where(&NodeInfo{Id: req.Id}).First(&rsp).Error; err != nil {
		return err
	}
	data := mydb.Model(&rsp).Updates(req)
	if data.RowsAffected <= 0 {
		return errors.New("updated nothing")
	}
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func (a *Node) NodeDelete(ctx context.Context, req *nodePt.NodeReq, rsp *nodePt.NodeRsp) error {
	log.Print("NodeDelete")
	var count int
	if err := mydb.Model(&NodeInfo{}).Where("parent_node = ?", req.Guid).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("has children,you should delete children first")
	}

	data := mydb.Where("id = ?", req.Guid).Delete(&nodePt.NodeInfo{})
	log.Print(data.RowsAffected)
	if data.RowsAffected <= 0 {
		return errors.New("deleted nothing")
	}
	if data.Error != nil {
		return data.Error
	}

	return nil
}

func main() {
	mydb, err = gorm.Open("mysql", "wz:Password123!@tcp(192.168.74.50:3306)/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer mydb.Close()
	mydb.SingularTable(true) //全局限制不是复数

	service := micro.NewService(
		micro.Name("go.micro.srv.v1.node"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	nodePt.RegisterNodeHandler(service.Server(), new(Node))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
