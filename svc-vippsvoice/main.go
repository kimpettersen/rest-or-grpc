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

	method := flag.String("action", "", "")
	amount := flag.String("amount", "", "")
	id := flag.String("id", "", "")

	flag.Parse()

	switch *method {
	case "PAY":
		pay(client, *amount)
	case "CONFIRM":
		confirm(client, *id)
	case "GET":
		getById(client, *id)
	case "ALL":
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

func confirm(client pb.PaymentsClient, id string) {
	rsp, _ := client.Confirm(context.Background(), &pb.PaymentByIdRequest{
		Id: id,
	})
	fmt.Printf("Response: %v\n", rsp)
}

func getById(client pb.PaymentsClient, id string) {
	rsp, _ := client.GetById(context.Background(), &pb.PaymentByIdRequest{
		Id: id,
	})
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
