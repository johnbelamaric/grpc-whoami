package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/johnbelamaric/grpc-whoami/certs"
	"github.com/johnbelamaric/grpc-whoami/pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

)

type server struct {
	name string
}

func (s *server) Whoami(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	p, ok := peer.FromContext(ctx);
	if !ok {
		return nil, fmt.Errorf("Could not find peer in gRPC context.")
	}
	issuer := "n/a"
	subject := "n/a"
	if t, ok := p.AuthInfo.(credentials.TLSInfo); ok {
		if t.State.VerifiedChains != nil && len(t.State.VerifiedChains) > 0 {
			if t.State.VerifiedChains[0] != nil && len(t.State.VerifiedChains[0]) > 0 {
				issuer = t.State.VerifiedChains[0][0].Issuer.CommonName
				subject = t.State.VerifiedChains[0][0].Subject.CommonName
			}
		}
	}
	log.Printf("%s %s %s\n", p.Addr.String(), issuer, subject)
	resp := &pb.Response{ServerName: s.name, ClientIp: p.Addr.String(), ClientIssuer: issuer, ClientSubject: subject}
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
		log.Printf("Started on %s with server name '%s'.\n", serviceEP, serverName)
	}
	tlsConfig, err := certs.NewServerTLSConfig(cert, key, ca)
	if err != nil {
		panic(err)
	}
	service, err := net.Listen("tcp", serviceEP)
	if err != nil {
		panic(err)
	}

	var s *grpc.Server
	s = grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	pb.RegisterWhoamiServer(s, whoami)
	s.Serve(service)
}
