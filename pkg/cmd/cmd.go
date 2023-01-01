package cmd

import (
	"fmt"
	"github.com/f-robo/pb-api-gateway/pkg/auth/pb"
	"github.com/f-robo/pb-auth-service/pkg/configuration"
	"github.com/f-robo/pb-auth-service/pkg/service"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"time"
)

type Cmd interface {
	Run() error
	String() string
}
type command struct {
	service bool
}

func (c *command) String() string {
	return "pb-api-gateway"
}
func (c *command) Run() error {
	go func() {
		lis, err := net.Listen("tcp", configuration.RuntimeConf.Grpc.AuthSvcUrl)
		if err != nil {
			fmt.Println("Faild to listing:", err)
		}
		fmt.Println("auth grpc run")
		grpcServer := grpc.NewServer()
		s := &service.Server{}
		pb.RegisterAuthServiceServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("Failed to serve:", err)
		}
	}()

	return nil
}

var (
	DefaultCmd Cmd = New()
)

func New() *command {

	cmd := new(command)
	return cmd
}

func init() {
	rand.Seed(time.Now().Unix())
}
func Run() {
	if err := DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
