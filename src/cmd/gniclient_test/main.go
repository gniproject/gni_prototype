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

type info [][]byte

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

	address := []byte("localhost:10161")
	target := []byte("Test-onos-config")
	caPath := []byte("/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/onfca.crt")
	certPath := []byte("/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.crt")
	keyPath := []byte("/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.key")
	timeOut := []byte("10")

	infoTemp := info{address, target, caPath, certPath, keyPath, timeOut}

	// Test FetchRequst function based on gNMI Capabilities and Get RPCs.
	request := ""
	capbilityRequest := &gpb.CapabilityRequest{}
	reqProto := &request
	if err := proto.UnmarshalText(*reqProto, capbilityRequest); err != nil {
		fmt.Errorf("unable to parse gnmi.CapabilityRequest from %q : %v", *reqProto, err)
	}
	gnmiCapabilities := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiCapabilityRequest{capbilityRequest}, Metadata: infoTemp}

	response, err := c.Fetch(ctx, gnmiCapabilities)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Recevied gNMI Capabilities Response: %s", response.GetGnmiCapabilityResponse())

	request = "path: <elem: <name: 'system'> elem:<name:'config'> elem: <name: 'hostname'>>"
	getRequest := &gpb.GetRequest{}
	reqProto = &request
	if err = proto.UnmarshalText(*reqProto, getRequest); err != nil {
		fmt.Errorf("unable to parse gnmi.GetRequest from %q : %v", *reqProto, err)
	}

	gnmiGetReq := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiGetRequest{getRequest}, Metadata: infoTemp}

	response, err = c.Fetch(ctx, gnmiGetReq)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Recevied gNMI GetRequest Response: %s", response.GetGnmiGetResponse())
}
