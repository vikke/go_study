package main

import (
	"context"
	"fmt"
	"net"

	pb "fabric-tokyo-hermes/idl"
	"google.golang.org/grpc"
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
