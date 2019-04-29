package gniserver

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	southbound "github.com/gniproject/gni_prototype/src/pkg/southbound"
)

// GniServer gNI Server Struct
type GniServer struct{}

// SetTargetInfo creates a southbound device and returns a target based on
// the given info.
func SetTargetInfo(info [][]byte) (southbound.Device, southbound.Target) {
	// Parse the given metadata for creating a new device and target.
	address := bytes.NewBuffer(info[0]).String()
	targetInput := bytes.NewBuffer(info[1]).String()
	caPath := bytes.NewBuffer(info[2]).String()
	certPath := bytes.NewBuffer(info[3]).String()
	keyPath := bytes.NewBuffer(info[4]).String()
	timeOut, _ := time.ParseDuration((bytes.NewBuffer(info[5]).String()))
	// Create a device.
	device := southbound.Device{
		Addr:     address,
		Target:   targetInput,
		CaPath:   caPath,
		CertPath: certPath,
		KeyPath:  keyPath,
		Timeout:  timeOut,
	}

	// Return the target based on the address of the device.
	target, err := southbound.GetTarget(southbound.Key{Key: device.Addr})
	if err != nil {
		fmt.Println("Creating device for addr: ", device.Addr)
		target, _, err = southbound.ConnectTarget(device)
		if err != nil {
			fmt.Println("Error ", target, err)
		}
	}

	return device, target

}

// Fetch implements gNI Fetch RPC function that allows to read one
// or more P4 entities or retrieve a snapshot of data from
// the target using gNMI.
func (s *GniServer) Fetch(ctx context.Context, req *gni.FetchRequest) (*gni.FetchResponse, error) {

	info := req.GetMetadata()
	var err error
	_, target := SetTargetInfo(info)

	fetchResponse := &gni.FetchResponse{}
	switch req.Frequest.(type) {
	// Process a gNMI GetRquest
	case *gni.FetchRequest_GnmiGetRequest:
		log.Println("Recevied gNMI GetRequest")
		getResponse, getErr := southbound.Get(target, req.GetGnmiGetRequest())
		if getErr != nil {
			fmt.Println("Error ", target, err)
		}
		fetchResponse = &gni.FetchResponse{Fresponse: &gni.FetchResponse_GnmiGetResponse{getResponse}}
		break
		// Process a gNMI CapabilityRequest
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

	info := req.GetMetadata()
	var err error
	_, target := SetTargetInfo(info)

	return &gni.StoreResponse{}, nil
}

func (s *GniServer) Command(ctx context.Context, req *gni.CommandRequest) (*gni.CommandResponse, error) {

	return &gni.CommandResponse{}, nil
}

func (s *GniServer) Stream(streamServer gni.GNI_StreamServer) error {

	return nil
}
