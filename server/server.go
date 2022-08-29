package main

import (
	"context"
	"fmt"
	redisTools "keycluster/internal/redis"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "keycluster/proto"
)

type server struct{}

var data = make(map[string]string)

func (s server) Put(ctx context.Context, request *pb.PutRequest) (*pb.PutResponse, error) {
	key := request.GetKey()
	value := request.GetValue()
	data[key] = value
	return &pb.PutResponse{Resp: "Successfull"}, nil
}

func (s server) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	key := request.GetKey()
	value, exists := data[key]
	if exists {
		return &pb.GetResponse{Value: value}, nil
	}
	return &pb.GetResponse{Value: ""}, fmt.Errorf("key %s not found", key)
}

func main() {
	const PORT = 50051
	var nodeID = redisTools.RegisterNode()
	data["node_id"] = strconv.Itoa(nodeID)

	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("error on listenning: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterBrokerServer(s, &server{})
	log.Println("starting server on port " + strconv.Itoa(PORT))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("error on serving: %v\n", err)
	}
}
