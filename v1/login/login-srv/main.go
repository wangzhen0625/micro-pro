package main

import (
	"log"
	"time"

	// "fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/liudng/godump"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/etcdv3"
	"golang.org/x/net/context"
	proto "micro-pro/v1/proto/login"
	// "time"
	// "reflect"
)

type AdNode struct {
	Id     uint
	Name   string
	Title  string
	Child  uint
	Status uint
	Remark string
	Sort   uint
	Pid    uint
	Level  uint
	// CreatedAt time.Time  `gorm:"column:create_time"`
	// UpdatedAt time.Time  `gorm:"column:update_time"`
	// DeletedAt *time.Time `gorm:"column:delete_time"`
}

var redisConn redis.Conn
var mysqlDb *gorm.DB

type Login struct{}

func (l *Login) UserPwdLogin(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Login.Hello request")

	rsp.Msg = "Hello " + req.Name
	return nil
}
func (l *Login) CheckAuth(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Login.Hello request")

	rsp.Msg = "Hello " + req.Name
	return nil
}
func (l *Login) GetAdminInfo(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Login.Hello request")

	v, err := redisConn.Do("SET", "name", "red")
	if err != nil {
		rsp.Msg = "1"
		return nil
	}
	v, err = redis.String(redisConn.Do("GET", "name"))
	if err != nil {
		rsp.Msg = "2"
		return nil
	}
	godump.Dump(v)

	rsp.Msg = "Hello2 " + req.Name

	var adnode []AdNode
	mysqlDb.Find(&adnode)
	godump.Dump(len(adnode))
	for k, v := range adnode {
		godump.Dump(k)
		godump.Dump(v)
	}

	return nil
}

func main() {
	var err error
	redisConn, err = redis.Dial("tcp", "192.168.74.84:6379")

	if err != nil {
		log.Fatal(err)
	}
	defer redisConn.Close()

	mysqlDb, err = gorm.Open("mysql", "root:P@ssw0rd1!@tcp(192.168.74.84:3306)/topsecgw?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer mysqlDb.Close()
	mysqlDb.SingularTable(true) //全局限制不是复数

	service := micro.NewService(
		micro.Name("go.micro.srv.v1.login"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	proto.RegisterLoginHandler(service.Server(), new(Login))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
