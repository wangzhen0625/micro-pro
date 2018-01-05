package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	accPt "micro-pro/v1/proto/account"
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
type Account struct{}

func (a *Account) NodeIndex(ctx context.Context, req *accPt.NodeReq, rsp *accPt.NodeRsp) error {
	log.Print("Received srv account.NodeIndex request")
	user := User{}

	rsp.Children = user.Children
	rsp.Id = user.Id
	rsp.Liattr.Guid = user.Liattr.Guid
	rsp.Liattr.Type = user.Liattr.Type
	rsp.Parent = user.Parent
	rsp.Text = user.Text
	rsp.Type = user.Type
	return nil
}

func main() {
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
