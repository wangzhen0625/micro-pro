# micro-pro
个人项目

在一个纯净的centos最小系统中，只要安装并运行etcd micro api srv二进制文件，
 即部署完了微服务，不需要任何的额外依赖。不需要go环境，也不需要micro源码和库。全部编译在了二进制文件中。

1、etcd 服务发现

//micro反向代理rest的统一入口，会解析url请求到 api
2、micro --registry=etcdv3 --registry_address=127.0.0.1:2379 api --handler=proxy

//所有的api，处理代码逻辑，会通过srv的client来调用srv来操作数据库，举例：
3、rest --registry=etcdv3 --registry_address=127.0.0.1:2379

//所有的srv，连接数据库返回数据，举例：
4、test-srv --registry=etcdv3 --registry_address=127.0.0.1:2379


protoc --go_out=plugins=micro:. hi.proto
定义RPC服务接口和对应的方法，go通过实现对应的方法来实现该接口
-------------------------------------------------------------
一个模块对应一个服务

service Login {...} 
=> 
结构体：Login 客户端接口：LoginClient   客户端新建方法：NewLoginClient  

11
-------------------------------------------------------------
非restful 接口(因为micro api使用的是反向代理的方式，所有这个接口用的也是go-restful，只是指定方法了)

0、认证授权（/v1/login/） 任何人都可以访问

4、系统管理（/v1/system/）
5、网络管理（/v1/network/）
6、PKI管理（/v1/pki/）
8、日志管理（/v1/log/）


restful接口

1、用户管理（/v1/user/）
	/v1/user/accounts
			/v1/user/accounts/pid/certs
			/v1/user/accounts/pid/passcodes
	/v1/user/attributes
	/v1/user/applys
	/v1/user/roles
	/v1/user/onlines
2、认证管理（/v1/authentication/）
	/v1/authentication/authenticationmethods
	/v1/authentication/passwordpolicys
	/v1/authentication/loginpagemanagerments
	/v1/authentication/portalpagemanagerments
3、应用管理(/v1/application/）
	/v1/application/applications
	/v1/application/secondaryaccounts
	/v1/application/ssolists
	/v1/application/ssofiles
	/v1/application/appimgs
	/v1/application/rules
7、消息管理（/v1/message/）
	/v1/message/msgways
	/v1/message/mailtemplates
	/v1/message/mailservices
9、图表管理（/v1/chart/）
	/v1/chart/charts







