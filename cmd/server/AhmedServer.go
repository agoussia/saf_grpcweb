// AhmedServer project main.go

/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	//	"crypto/tls"
	"flag"
	"log"
	//	"net"
	"fmt"

	"github.com/xid"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	dir "github.com/saf_grpcweb/gen"

	// "google.golang.org/grpc/reflection"

	"github.com/grpc-web/go/grpcweb"

	"net/http"

	"google.golang.org/grpc/grpclog"
)

// server is used to implement ui.BuyerService
type server struct{}

func (s *server) GetPersonInfo(ctx context.Context, in *dir.Ack) (*dir.Person, error) {

	person := dir.Person{

		FirstName: "Alex",
		LastName:  "Goussiatiner",
		Address: &dir.Address{
			Street: "abcd street",
			City:   "Doha",
		},
	}

	return &person, nil

}

func (s *server) SetPersonInfo(ctx context.Context, in *dir.Person) (*dir.Ack, error) {

	log.Printf("++++++Person: %v", in.String())

	return &dir.Ack{}, nil

}

var (
	//http1Port       = flag.Int("http1_port", 9090, "Port to listen with HTTP1.1 with TLS on.")
	//http1EmptyPort  = flag.Int("http1_empty_port", 9095, "Port to listen with HTTP1.1 with TLS on with a grpc server that has no services.")
	http2Port = flag.Int("http2_port", 9100, "Port to listen with HTTP2 with TLS on.")
	//http2EmptyPort  = flag.Int("http2_empty_port", 9105, "Port to listen with HTTP2 with TLS on with a grpc server that has no services.")
	tlsCertFilePath = flag.String("tls_cert_file", "misc/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "misc/localhost.key", "Path to the private key file.")
)

func main() {

	guid := xid.New()

	println(guid.String())

	log.Printf("Machine: %v", guid.Machine())

	log.Printf("Pid: %v", guid.Pid())

	log.Printf("Time: %v", guid.Time())

	log.Printf("Counter: %v", guid.Counter())

	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	grpcServer := grpc.NewServer()

	/*
		dir.RegisterBuyerServiceServer(s, &server{})

		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	*/

	//wrappedGrpc := grpcweb.WrapServer(grpcServer)
	//tlsHttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

	websocketOriginFunc := grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	})

	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true), websocketOriginFunc)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	http2Server := http.Server{
		Addr:    fmt.Sprintf(":%d", *http2Port),
		Handler: http.HandlerFunc(handler),
	}

	/*
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
		}
		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)

	*/

	// Start the Http2 server
	if err := http2Server.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}

}
