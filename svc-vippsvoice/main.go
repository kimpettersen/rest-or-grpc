package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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

	action := flag.String("action", "", "")
	amount := flag.String("amount", "", "")
	flag.Parse()

	switch *action {
	case "pay":
		pay(client, *amount)
	case "getall":
		getAll(client)
	default:
		fmt.Println("unknown argument")
	}
}

func pay(client pb.PaymentsClient, amount string) {
	a, _ := strconv.ParseInt(amount, 10, 64)
	payment := &pb.PaymentRequest{
		Amount: a,
		From:   "1234321",
		To:     "234432",
	}
	rsp, _ := client.Pay(context.Background(), payment)
	fmt.Printf("Response: %v\n", rsp)
}

func getAll(client pb.PaymentsClient) {
	rsp, _ := client.GetAll(context.Background(), &pb.AllPaymentsRequest{})

	for idx, p := range rsp.Payments {
		s := fmt.Sprintf("Payment #%v: %v\n", idx, p)
		fmt.Println(strings.Repeat("-", len(s)))
		fmt.Print(s)
	}
}
