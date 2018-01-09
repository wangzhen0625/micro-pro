package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"log"
	accPt "micro-pro/v1/proto/account"
	"time"
)

type Account struct{}

var (
	mydb *gorm.DB
	err  error
)

type Role struct {
	Id        uint
	Name      string
	Desc      string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	/*CreatedAt  time.Time  `gorm:"column:create_time"`
	UpdatedAt  time.Time  `gorm:"column:update_time"`
	DeletedAt  *time.Time `gorm:"column:delete_time"`*/
}

func (a *Account) NodeIndex(ctx context.Context, req *accPt.NodeReq, rsp *accPt.NodeRsp) error {
	log.Print("Received srv account.NodeIndex request")
	var node NodeInfo
	mydb.Where(&NodeInfo{ParentNode: req.ParentId}).Select("id,parent_node,type,login_id").First(&node)
	log.Print(node)

	//  {children: true, id: "3", li_attr:{guid:"3",type: "organization"}, parent: "#", text: "topsec",type: "organization"}
	liattr := accPt.LiAttr{}
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
func (a *Account) NodeChildren(ctx context.Context, req *accPt.NodeReq, stream accPt.Account_NodeChildrenStream) error {
	var nodes []NodeInfo
	mydb.Where(&NodeInfo{ParentNode: req.ParentId}).Select("id,parent_node,type,login_id").Find(&nodes)

	feature := accPt.NodeRsp{}
	liattr := accPt.LiAttr{}
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

func main() {
	mydb, err = gorm.Open("mysql", "wz:Password123!@tcp(192.168.74.50:3306)/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer mydb.Close()
	mydb.SingularTable(true) //全局限制不是复数

	service := micro.NewService(
		micro.Name("go.micro.srv.v1.account"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	accPt.RegisterAccountHandler(service.Server(), new(Account))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
