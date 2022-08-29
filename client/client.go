package main

import (
	"context"
	"fmt"
	pb "keycluster/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	url := "192.168.59.102:30005"
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	c = pb.NewBrokerClient(conn)

	val, _ := get("node_id")
	fmt.Println(val)
}
