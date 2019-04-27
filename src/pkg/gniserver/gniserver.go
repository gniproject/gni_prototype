package gniserver

import (
	"fmt"
	"log"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	southbound "github.com/gniproject/gni_prototype/src/pkg/southbound"

	"context"
)

const (
	port = ":50051"
)

type GniServer struct{}

func createDevice() southbound.Device {
	device := southbound.Device{
		Addr:     "localhost:10161",
		Target:   "Test-onos-config",
		CaPath:   "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/onfca.crt",
		CertPath: "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.crt",
		KeyPath:  "/home/adib/gni_prototype/src/github.com/opennetworkinglab/onos-config/tools/test/devicesim/certs/client1.key",
		Timeout:  10,
	}

	return device
}

func getTarget(device southbound.Device) (southbound.Target, error) {

	target, err := southbound.GetTarget(southbound.Key{Key: device.Addr})
	if err != nil {
		fmt.Println("Creating device for addr: ", device.Addr)
		target, _, err = southbound.ConnectTarget(device)
		if err != nil {
			fmt.Println("Error ", target, err)
		}
	}
	return target, err
}

func (s *GniServer) Fetch(ctx context.Context, req *gni.FetchRequest) (*gni.FetchResponse, error) {

	device := createDevice()
	target, err := getTarget(device)

	fetchResponse := &gni.FetchResponse{}
	switch req.Frequest.(type) {
	case *gni.FetchRequest_GnmiGetRequest:
		log.Println("Recevied gNMI GetRequest")
		getResponse, getErr := southbound.Get(target, req.GetGnmiGetRequest())
		if getErr != nil {
			fmt.Println("Error ", target, err)
		}
		fetchResponse = &gni.FetchResponse{Fresponse: &gni.FetchResponse_GnmiGetResponse{getResponse}}
		break

	case *gni.FetchRequest_GnmiCapabilityRequest:
		log.Println("Recevied gNMI Capabilityrequest")
		getResponse, getErr := southbound.Capabilities(target, req.GetGnmiCapabilityRequest())
		if getErr != nil {
			fmt.Println("Error ", target, err)
		}
		fetchResponse = &gni.FetchResponse{Fresponse: &gni.FetchResponse_GnmiCapabilityResponse{getResponse}}
		break

	}
	return fetchResponse, nil
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
