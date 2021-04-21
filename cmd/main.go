package main 

import (
	"fmt"
	"net"
	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"github.com/JIeeiroSst/go-app/proto"
	"github.com/JIeeiroSst/go-app/log"
	grpcConfig"github.com/JIeeiroSst/go-app/delivery/grpc"
	"google.golang.org/grpc"
)

func main(){
	mysqlConn := mysql.NewMysqlConn(&config.Config.MysqlConfig)
	service:=grpcConfig.NewService(mysqlConn)
	lis, err := net.Listen("tcp", config.Config.PORT)
	if err != nil {
		log.InitZapLog().Error(fmt.Sprintf("failed to listen: %v", err))
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserProfileServer(grpcServer,service.UnimplementedUserProfileServer)
	log.InitZapLog().Info(fmt.Sprintf("server have localhost:" + config.Config.PORT + " is running"))
	grpcServer.Serve(lis)
}