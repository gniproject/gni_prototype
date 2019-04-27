package main

import (
	"context"
	"fmt"
	"log"
	"time"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	"github.com/golang/protobuf/proto"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
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

	// Test FetchRequst function based on gNMI Capabilities and Get RPCs.
	request := ""
	capbilityRequest := &gpb.CapabilityRequest{}
	reqProto := &request
	if err := proto.UnmarshalText(*reqProto, capbilityRequest); err != nil {
		fmt.Errorf("unable to parse gnmi.CapabilityRequest from %q : %v", *reqProto, err)
	}
	gnmiCapabilities := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiCapabilityRequest{capbilityRequest}}

	response, err := c.Fetch(ctx, gnmiCapabilities)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Recevied gNMI Capabilities Response: %s", response.GetGnmiCapabilityResponse())

	request = "path: <elem: <name: 'system'> elem:<name:'config'> elem: <name: 'hostname'>>"
	getRequest := &gpb.GetRequest{}
	reqProto = &request
	if err = proto.UnmarshalText(*reqProto, getRequest); err != nil {
		fmt.Errorf("unable to parse gnmi.GetRequest from %q : %v", *reqProto, err)
	}

	gnmiGetReq := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiGetRequest{getRequest}}

	response, err = c.Fetch(ctx, gnmiGetReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Recevied gNMI GetRequest Response: %s", response.GetGnmiGetResponse())
}
