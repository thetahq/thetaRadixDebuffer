package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"thetaRadixDebuffer/thetaradix"
)

func main() {
	log.Println("Starting Theta Radix® Debufffer")
	log.Println("Purpose of this software is connecting gRPC backend with REST based")
	startGrpcServer()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:8557") // TODO: Use config file
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	thetaradix.RegisterThetaRadixServer(grpcServer, DebufferServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

type DebufferServer struct{}

func (DebufferServer) Register(context.Context, *thetaradix.RegisterRequest) (*thetaradix.RegisterReply, error) {
	panic("implement me")
}
