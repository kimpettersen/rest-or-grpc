package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/kimpettersen/svc-payments/proto"
	"google.golang.org/grpc"
)

func main() {
	address := "127.0.0.1:3000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %+v", err)
	}
	defer conn.Close()
	client := pb.NewPaymentsClient(conn)

	id := flag.String("id", "", "specify id of payment")
	flag.Parse()

	p, err := client.ConfirmPayment(context.Background(), &pb.PaymentId{
		Id: *id,
	})

	fmt.Println(p.GetStatus())

}
