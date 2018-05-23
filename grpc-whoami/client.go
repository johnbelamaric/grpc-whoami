package main

import (
	"flag"
	"log"

	"github.com/johnbelamaric/grpc-whoami/certs"
	"github.com/johnbelamaric/grpc-whoami/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	creds "google.golang.org/grpc/credentials"

)

func main() {
	var (
	endpoint string
	verbose bool
                cert string
                key string
                ca string
	)

	flag.BoolVar(&verbose, "v", false, "Log verbosely")
        flag.StringVar(&cert, "cert", "", "TLS cert PEM file path")
        flag.StringVar(&key, "key", "", "TLS key PEM file path")
        flag.StringVar(&ca, "ca", "", "TLS ca cert PEM file path")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		endpoint = args[0]
	} else {
		endpoint = "localhost:8123"
	}

	if verbose {
		log.Printf("endpoint: %s\n", endpoint)
	}

	tlsConfig, err := certs.NewTLSConfig(cert, key, ca)
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(creds.NewTLS(tlsConfig)))
	client := pb.NewWhoamiClient(conn)

	resp, err := client.Whoami(context.Background(), &pb.Request{})

	if err == nil {
		log.Printf("%s %s %s %s\n", resp.ServerName, resp.ClientIp, resp.ClientIssuer, resp.ClientSubject)
	} else {
		log.Printf("error: %s\n", err)
	}
}
