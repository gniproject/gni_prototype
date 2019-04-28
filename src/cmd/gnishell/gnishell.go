package main

import (
	"context"
	"flag"
	"log"
	"time"

	gni "github.com/gniproject/gni_prototype/src/api/gni"
	"google.golang.org/grpc"
)

var (
	fetch         = flag.String("fetch", "", "When set, CLI will perform a Fetch request..")
	reqProto      = flag.String("proto", "", "Text proto for requests.")
	targetAddress = flag.String("targetaddress", "", "Address of a target in the format of ip:port")
	gniAddress    = flag.String("gniaddress", "", "Address of a gNI server in the format of ip:port")
	timeoOut      = flag.Int("timeout", 10, "Terminate query if no RPC is established within the timeout duration.")

	// Certificate files.
	caCert     = flag.String("ca_crt", "", "CA certificate file. Used to verify server TLS certificate.")
	clientCert = flag.String("client_crt", "", "Client certificate file. Used for client certificate-based authentication.")
	clientKey  = flag.String("client_key", "", "Client private key file. Used for client certificate-based authentication.")
)

type info [][]byte

func main() {

	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*gniAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gni.NewGNIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

}
