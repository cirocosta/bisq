package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/bisq-network/bisq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:9998"
	defaultName = "world"
)

func run(ctx context.Context) error {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	defer conn.Close()

	walletsClient := pb.NewWalletsClient(conn)

	ctx = metadata.AppendToOutgoingContext(ctx, "password", "abctesting123")
	balance, err := walletsClient.GetAddressBalance(ctx, &pb.GetAddressBalanceRequest{
		Address: "bcrt1q6j90vywv8x7eyevcnn2tn2wrlg3vsjlsvt46qz",
	})
	if err != nil {
		return fmt.Errorf("get address balance: %w", err)
	}

	fmt.Printf("balance=%+v\n", balance)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Panic(err)
	}
}
