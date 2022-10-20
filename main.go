package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jainachal03/simple/proto"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

const (
	port    = ":9000"
	address = "localhost:9000"
)

type Server struct {
	proto.UnimplementedSimpleServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println(" the server is listening on port : 9000")
	grpcServer := grpc.NewServer()
	proto.RegisterSimpleServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *Server) GetSimple(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleReply, error) {
	res := &proto.SimpleReply{}
	if req == nil {
		fmt.Println("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}
	if req.Request == "" {
		fmt.Println("name must not be empty in the request")
		return res, xerrors.Errorf("name must not be empty in the request")
	}
	log.Printf("Received: %v", req.GetRequest())
	response := "This is the default response :-)"
	res.Reply = response
	return res, nil

}
