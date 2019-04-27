package gniserver

import (
	"fmt"
	"log"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	southbound "github.com/gniproject/gni_prototype/src/pkg/southbound"
	"github.com/golang/protobuf/proto"

	"context"
)

const (
	port = ":50051"
)

type GniServer struct{}

func (s *GniServer) Fetch(ctx context.Context, req *gni.FetchRequest) (*gni.FetchResponse, error) {

	device := southbound.Device{
		Addr:     "localhost:10161",
		Target:   "Test-onos-config",
		CaPath:   "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/onfca.crt",
		CertPath: "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.crt",
		KeyPath:  "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.key",
		Timeout:  10,
	}
	target, err := southbound.GetTarget(southbound.Key{Key: device.Addr})
	if err != nil {
		fmt.Println("Creating device for addr: ", device.Addr)
		target, _, err = southbound.ConnectTarget(device)
		if err != nil {
			fmt.Println("Error ", target, err)
		}
	}

	request := ""
	capResponse, capErr := southbound.CapabilitiesWithString(target, request)
	if capErr != nil {
		fmt.Println("Error ", target, err)
	}
	capResponseString := proto.MarshalTextString(capResponse)
	fmt.Println("Capabilities: ", capResponseString)
	request = "path: <elem: <name: 'system'> elem:<name:'config'> elem: <name: 'hostname'>>"
	getResponse, getErr := southbound.GetWithString(target, request)
	if getErr != nil {
		fmt.Println("Error ", target, err)
	}
	getResponseString := proto.MarshalTextString(getResponse)
	fmt.Println("get: ", getResponseString)

	log.Println("Fetch Request is arrived")
	switch req.Frequest.(type) {
	case *gni.FetchRequest_GnmiGetRequest:
		log.Println("Gnmi Get Request")
		break
	case *gni.FetchRequest_GnmiCapabilityRequest:
		log.Println("Gnmi Capability Request")
		break

	}
	return &gni.FetchResponse{}, nil
}

func (s *GniServer) Store(ctx context.Context, req *gni.StoreRequest) (*gni.StoreResponse, error) {

	return &gni.StoreResponse{}, nil
}

func (s *GniServer) Command(ctx context.Context, req *gni.CommandRequest) (*gni.CommandResponse, error) {

	return &gni.CommandResponse{}, nil
}

func (s *GniServer) Stream(streamServer gni.GNI_StreamServer) error {

	return nil
}
