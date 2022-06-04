package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "restful/restful"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:7501", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRestfulClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &pb.GetRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}
	fmt.Printf("The task's name is \"%s\" and its status is %s\n", r.GetTitle(), r.GetStatus())
}
