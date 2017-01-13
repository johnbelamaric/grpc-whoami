package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/johnbelamaric/grpc-whoami/certs"
	"github.com/johnbelamaric/grpc-whoami/pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type server struct {
	name string
}

func (s *server) Whoami(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{ServerName: s.name}
	return resp, nil
}

func main() {
	var (
		verbose    bool
		serviceEP  string
		serverName string
		cert string
		key string
		ca string
	)

	flag.BoolVar(&verbose, "v", false, "Log verbosely")
	flag.StringVar(&serviceEP, "l", "0.0.0.0:8123", "listen on this address:port")
	flag.StringVar(&serverName, "n", "", "server name to use")
	flag.StringVar(&cert, "cert", "", "TLS cert PEM file path")
	flag.StringVar(&key, "key", "", "TLS key PEM file path")
	flag.StringVar(&ca, "ca", "", "TLS ca cert PEM file path")
	flag.Parse()

	if serverName == "" {
		serverName = os.Getenv("HOSTNAME")
	}

	if serverName == "" {
		name, err := os.Hostname()
		if err != nil {
			serverName = err.Error()
		} else {
			serverName = name
		}
	}

	whoami := &server{name: serverName}

	if verbose {
		fmt.Printf("Started on %s with server name '%s'.\n", serviceEP, serverName)
	}
	tlsConfig, err := certs.NewTLSConfig(cert, key, ca)
	if err != nil {
		panic(err)
	}
	service, err := tls.Listen("tcp", serviceEP, tlsConfig)
	if err != nil {
		panic(err)
	}

	var s *grpc.Server
	s = grpc.NewServer()

	pb.RegisterWhoamiServer(s, whoami)
	s.Serve(service)
}
