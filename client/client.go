package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "keycluster/proto"
	"log"
)

var c pb.BrokerClient

func put(key, value string) {
	_, err := c.Put(context.Background(), &pb.PutRequest{Key: key, Value: value})
	if err != nil {
		log.Fatalf("could not put: %v\n", err)
	}
}
func get(key string) (string, error) {
	resp, err := c.Get(context.Background(), &pb.GetRequest{Key: key})
	return resp.GetValue(), err
}

func main() {
	url := "localhost:50051"
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	c = pb.NewBrokerClient(conn)

	put("key1", "value1")
	put("key2", "value2")
	get("key1")
	if _, err := get("key3"); err == nil {
		log.Fatalf("key3 should not exist\n")
	}
}
