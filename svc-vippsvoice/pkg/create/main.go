package main

import (
	"context"
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
	payment := &pb.Payment{
		Amount: 1000,
		From:   "1234321",
		To:     "234432",
	}
	rsp, err := client.Pay(context.Background(), payment)
	fmt.Printf("\n----------> ERR: %v\n\n", err)
	fmt.Printf("\n----------> ID: %v\n", rsp.Id)
	fmt.Printf("\n----------> AMOUNT: %v\n", rsp.Amount)
	fmt.Printf("\n----------> FROM: %v\n", rsp.From)
	fmt.Printf("\n----------> TO: %v\n", rsp.To)
	fmt.Printf("\n----------> STATUS: %v\n", rsp.GetStatus())

}
