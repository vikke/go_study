package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "hermes-pub/idl"
)

const (
	port = "50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	fmt.Println("aaa")
}
