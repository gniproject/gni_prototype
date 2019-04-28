package gniserver

import (
	"context"
	"fmt"
	"log"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	southbound "github.com/gniproject/gni_prototype/src/pkg/southbound"
)

type GniServer struct {
	Device southbound.Device
	Target southbound.Target
}

func NewServer(device southbound.Device, target southbound.Target) *GniServer {
	s := &GniServer{device, target}
	return s
}

type TargetInfo struct {
	Device southbound.Device
	Target southbound.Target
}

var TargetInfoChan chan southbound.Target

var info TargetInfo

func SetTargetInfo() TargetInfo {
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

	info = TargetInfo{device, target}

	log.Println("test" + info.Device.Addr)

	return info

}

// Fetch implements gNI Fetch RPC function that allows to read one
// or more P4 entities or retrieve a snapshot of data from
// the target using gNMI.
func (s *GniServer) Fetch(ctx context.Context, req *gni.FetchRequest) (*gni.FetchResponse, error) {

	log.Println("Fetch called")
	//target := info.Target
	target := s.Target

	log.Println(info.Device.Addr)
	var err error

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

// Store implements the gNI Store RPC function  that allows to
// update one or more P4 entities or modify the state of data
// on the target using gNMI.
func (s *GniServer) Store(ctx context.Context, req *gni.StoreRequest) (*gni.StoreResponse, error) {

	return &gni.StoreResponse{}, nil
}

func (s *GniServer) Command(ctx context.Context, req *gni.CommandRequest) (*gni.CommandResponse, error) {

	return &gni.CommandResponse{}, nil
}

func (s *GniServer) Stream(streamServer gni.GNI_StreamServer) error {

	return nil
}
