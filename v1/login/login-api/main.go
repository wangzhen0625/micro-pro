package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"github.com/liudng/godump"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-web"
	"golang.org/x/net/context"
	proto "micro-pro/v1/proto/login"
	"time"
)

type User struct {
	Name     string `json:"name" description:"name of the user" default:"john"`
	Password string `json:"password" description:"age of the user" default:"21"`
}

type Login struct{}

var cl proto.LoginClient

type MyCustomClaims struct {
	jwt.StandardClaims
	Type      string `json:"type"`
	LoginTime int64  `json:"logintime"`
	LoginIp   string `json:"loginip"`
}

func (l *Login) UserPwdLogin(req *restful.Request, rsp *restful.Response) {

	now := time.Now().Unix()
	end := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := MyCustomClaims{}

	claims.Audience = "name" //aud接收者
	claims.ExpiresAt = end   //exp生命周期
	claims.Id = "16"         //jti 用户ID
	claims.IssuedAt = now    //iat token创建时间
	claims.Issuer = "topsec" //iss jwt的签发者
	claims.NotBefore = now   //nbf  开始生效的时间
	claims.Subject = "name"  //sub JWT所面向的用户

	claims.Type = "4"
	claims.LoginTime = now
	claims.LoginIp = "192.168.74.50"

	SecretKey := "TOPSEC_GW_2017"
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte(SecretKey))

	// usr := new(User)
	// req.ReadEntity(&usr)

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	godump.Dump("------------1--------------------------------------------")
	godump.Dump(token)
	godump.Dump("--------------2------------------------------------------")
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		godump.Dump(claims)
		//解析成功
		godump.Dump("-----------3---------------------------------------------")
	} else {
		godump.Dump(err)
		godump.Dump("------------4--------------------------------------------")
	}

	// list := User{Name: "wz", Password: "12"}
	rsp.WriteEntity(tokenString)
}
func (l *Login) GetAdminInfo(req *restful.Request, rsp *restful.Response) {
	response, err := cl.GetAdminInfo(context.TODO(), &proto.Request{
		Name:     "name",
		Password: "password",
	})

	if err != nil {
		rsp.WriteError(500, err)
	}
	rsp.WriteEntity(response)

}
func (l Login) CreateContainer() *restful.Container {
	wc := restful.NewContainer()

	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path("/v1/login")
	ws.Route(ws.POST("/userPwdLogin").To(l.UserPwdLogin))
	ws.Route(ws.GET("/getAdminInfo").To(l.GetAdminInfo))
	wc.Add(ws)

	return wc
}

/*func (s *Say) Hello(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Say.Hello API request")

	name := req.PathParameter("name")

	response, err := cl.Hello(context.TODO(), &login.Request{
		Name: name,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}*/

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.v1.login"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)

	service.Init()

	// setup Greeter Server Client
	cl = proto.NewLoginClient("go.micro.srv.v1.login", client.DefaultClient)

	// Create RESTful handler
	login := new(Login)
	wc := login.CreateContainer()
	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
