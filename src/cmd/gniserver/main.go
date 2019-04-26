package main

import (
	"log"
	"net"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	gniserver "github.com/gniproject/gni_prototype/src/pkg/gniserver"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gni.RegisterGNIServer(grpcServer, &gniserver.GniServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
