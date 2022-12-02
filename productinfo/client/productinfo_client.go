package main

import (
	"context"
	"log"
	"time"

	pb "github.com/abdulmajid18/grpc/productinfo/service/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple iphone 11"
	description := `Meet Apple iphone 11 All-new dual-camera system with
	Ultra Wide and Night mode.`

	// price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx,
		&pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
