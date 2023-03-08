package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPC interface {
	Start(debug bool, conf GrpcConf) error
	Stop()
}

var (
	instance   *server = &server{}
	grpcServer *grpc.Server
)

func GetInstance( /*TODO: inject dependencies*/ ) GRPC {
	return instance
}

type server struct {
	/*TODO: servers or grpc services*/
}

type GrpcConf struct {
	ConnectionType string
	Address        string
}

func (s *server) Start(debug bool, conf GrpcConf) error {
	fmt.Println("* GRPC SERVER STARTING *")

	lis, err := s.makeListener(conf.ConnectionType, conf.Address)
	if err != nil {
		panic(err)
	}

	grpcServer = s.makeServer()

	err = s.registerServices(grpcServer)
	if err != nil {
		panic(err)
	}

	if debug {
		fmt.Println("reflection registered")
		reflection.Register(grpcServer)
	}

	s.serve(grpcServer, lis)
	return nil
}

func (s *server) Stop() {
	if grpcServer != nil {
		fmt.Println("* STOPPING GRPC SERVER *")
		grpcServer.Stop()
	}
}

func (s *server) makeListener(connectionType, address string) (net.Listener, error) {
	lis, err := net.Listen(connectionType, address)
	if err != nil {
		panic("implement error returning")
	}
	return lis, nil
}

func (s *server) makeServer() *grpc.Server {
  // Add interceptors from (montana.fedora.com) 
	// grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	// grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	return grpc.NewServer(
	/*
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcValidator.UnaryServerInterceptor(),
			),
		),
	*/
	)
}

func (s *server) registerServices(server *grpc.Server) error {
	/* register services */
	return nil
}

func (s *server) serve(server *grpc.Server, lis net.Listener) {
	go func() {
		if err := server.Serve(lis); err != nil {
			panic("failed to serve GRPC server")
		}
	}()
}
