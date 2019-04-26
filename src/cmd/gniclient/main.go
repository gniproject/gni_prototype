package main

import (
	"context"
	"log"
	"time"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	gpb "github.com/gniproject/gni_prototype/src/api/gnmi"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gni.NewGNIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	gnmiGetRequest := &gpb.GetRequest{Type: 1}

	gnmiGetReq := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiGetRequest{gnmiGetRequest}}

	r, err := c.Fetch(ctx, gnmiGetReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGnmiGetResponse())
}
