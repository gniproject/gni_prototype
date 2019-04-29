package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	gni "github.com/gniproject/gni_prototype/src/api/gni"

	"github.com/golang/protobuf/proto"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
)

// gni shell options and constants
var (
	rpc           = flag.String("rpc", "", "When set, CLI will perform one of the gNI RPCs (possible values: fetch, store, command, stream)")
	rpcType       = flag.String("rpc_type", "", "Specefiy the type of target and request in the format of TargetType:RquestType, e.g. gnmi:get")
	requestProto  = flag.String("proto", "", "Text proto for requests.")
	targetName    = flag.String("target_name", "", "Name of the Target")
	targetAddress = flag.String("target_address", "", "Address of a target in the format of ip:port")
	gniAddress    = flag.String("gni_address", "", "Address of a gNI server in the format of ip:port")
	timeout       = flag.String("timeout", "10", "Terminate query if no RPC is established within the timeout duration.")

	// Certificate files.
	caCert     = flag.String("ca_crt", "", "CA certificate file. Used to verify server TLS certificate.")
	clientCert = flag.String("client_crt", "", "Client certificate file. Used for client certificate-based authentication.")
	clientKey  = flag.String("client_key", "", "Client private key file. Used for client certificate-based authentication.")

	// rpc types
	GnmiGet        = "gnmi:get"
	GnmiCapability = "gnmi:capability"

	// rpc names
	Fetch   = "fetch"
	Store   = "store"
	Command = "command"
	Stream  = "stream"
)

type info [][]byte

func main() {

	flag.Parse()

	if gniAddress == nil {
		log.Fatalf("gNI server address is empty %s", *gniAddress)
		return
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(*gniAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Create a gNI client
	c := gni.NewGNIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// TODO check the options to make sure they are not null
	address := []byte(*targetAddress)
	target := []byte(*targetName)
	caCertPath := []byte(*caCert)
	clientCertPath := []byte(*clientCert)
	clientKeyPath := []byte(*clientKey)
	timeOut := []byte(*timeout)
	request := *requestProto

	infoTemp := info{address, target, caCertPath, clientCertPath, clientKeyPath, timeOut}

	// Test FetchRequst function based on gNMI Capabilities and Get RPCs.

	switch *rpc {
	case Fetch:
		switch *rpcType {
		case GnmiCapability:
			capbilityRequest := &gpb.CapabilityRequest{}
			reqProto := &request
			if err := proto.UnmarshalText(*reqProto, capbilityRequest); err != nil {
				fmt.Errorf("unable to parse gnmi.CapabilityRequest from %q : %v", *reqProto, err)
			}
			gnmiCapabilities := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiCapabilityRequest{capbilityRequest}, Metadata: infoTemp}

			response, err := c.Fetch(ctx, gnmiCapabilities)

			if err != nil {
				log.Fatalf("could not execute RPC on the remote server: %v", err)
			}
			fmt.Println("Recevied gNMI Capabilities Response:", response.GetGnmiCapabilityResponse())
			fmt.Println("----------------------------------------------------------------------------")
			break
		case GnmiGet:
			getRequest := &gpb.GetRequest{}
			reqProto := &request
			if err = proto.UnmarshalText(*reqProto, getRequest); err != nil {
				fmt.Errorf("unable to parse gnmi.GetRequest from %q : %v", *reqProto, err)
			}

			gnmiGetReq := &gni.FetchRequest{Frequest: &gni.FetchRequest_GnmiGetRequest{getRequest}, Metadata: infoTemp}

			response, err := c.Fetch(ctx, gnmiGetReq)
			if err != nil {
				log.Fatalf("could not execute RPC on the remote server: %v", err)
			}
			fmt.Println("Recevied gNMI GetRequest Response: ", response.GetGnmiGetResponse())
			fmt.Println("----------------------------------------------------------------------------")
			break
		}

	}

}
