package main

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/abdulmajid18/grpc/productinfo/service/ecommerce"
)

// server is used to implement ecommmerce/product_info

type server struct {
	productMap map[string]pb.Product
}

// AddProduct implements ecommerce.AddProduct

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generaating Product ID", error)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]pb.Product)
	}

}
