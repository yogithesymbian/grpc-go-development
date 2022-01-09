package main

import (
	"context"
	"fmt"
	"log"

	"grpc-go-development/greet/greetpb"

	"google.golang.org/grpc"
)

func main()  {
	fmt.Println("Hello i'm client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %v", err)
	}

	cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephane",
			LastName: "Maarek",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("response from Greet: %v", res.Result)

}