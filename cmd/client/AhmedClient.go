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
	//"fmt"
	"log"
	"os"
	"time"

	dir "github.com/saf_grpcweb/gen"

	"github.com/xid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	guid := xid.New()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := dir.NewBuyerServiceClient(conn)

	// Contact the server and print out its response.

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Name: %s", name)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	id := guid

	log.Printf("id: %v", id)

	ack := dir.Ack{}

	//
	r, err := c.GetPersonInfo(ctx, &ack)
	if err != nil {
		log.Fatalf("could not get Person: %v", err)
	}
	log.Printf("Person One: %s", r.GetFirstName())

	person := dir.Person{

		FirstName: "Baher",
		LastName:  "MAnsoor",
		Address: &dir.Address{
			Street: "abcd street",
			City:   "Doha",
		},
	}
	log.Printf("Person Two: %s", person.GetFirstName())
	_, err2 := c.SetPersonInfo(ctx, &person)
	if err2 != nil {
		log.Fatalf("could not get Person: %v", err2)
	}

}
