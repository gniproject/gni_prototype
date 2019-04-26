package gniserver

import (
	"context"
	"log"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
)

const (
	port = ":50051"
)

type GniServer struct{}

func (s *GniServer) Fetch(ctx context.Context, req *gni.FetchRequest) (*gni.FetchResponse, error) {

	log.Println("Test")

	return &gni.FetchResponse{}, nil
}

func (s *GniServer) Store(ctx context.Context, req *gni.StoreRequest) (*gni.StoreResponse, error) {

	log.Println("Test")

	return &gni.StoreResponse{}, nil
}

func (s *GniServer) Command(ctx context.Context, req *gni.CommandRequest) (*gni.CommandResponse, error) {

	log.Println("Test")

	return &gni.CommandResponse{}, nil
}

func (s *GniServer) Stream(streamServer gni.GNI_StreamServer) error {

	return nil
}
