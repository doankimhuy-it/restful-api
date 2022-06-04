package main

import (
	"context"
	pb "restful/restful"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGet(t *testing.T) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRestfulClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Get(ctx, &pb.GetRequest{Id: 1})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestCreate(t *testing.T) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRestfulClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Create(ctx, &pb.CreateRequest{Id: 100, Title: "Test", Status: "Done"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRestfulClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Update(ctx, &pb.UpdateRequest{Id: 100, Title: "Renamed", Status: "Doing"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestDelete(t *testing.T) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRestfulClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Delete(ctx, &pb.DeleteRequest{Id: 100})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
