package main

import (
	"flag"
	"log"

	"github.com/aweist/grpc_server/blog"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:8080", "The server address in the format of host:port")
)

func main() {
	var opts []grpc.DialOption
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial connection: %v", err)
	}
	defer conn.Close()

	blog.NewBlogServiceClient()
}
