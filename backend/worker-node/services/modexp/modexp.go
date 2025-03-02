package modexp

import (
	"math/big"
	"context"
	pb "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/modexp"
)

// ModExpServiceServer struct
type ModExpServiceServer struct {
	pb.UnimplementedModExpServiceServer
}

// ComputeModExp gRPC method
func (s *ModExpServiceServer) ComputeModExp(ctx context.Context, req *pb.ModExpRequest) (*pb.ModExpResponse, error) {
	base := new(big.Int)
	base.SetString(req.Base, 10)
	exponent := new(big.Int)
	exponent.SetString(req.Exponent, 10)
	modulus := new(big.Int)
	modulus.SetString(req.Modulus, 10)

	result := new(big.Int).Exp(base, exponent, modulus)
	return &pb.ModExpResponse{Result: result.String()}, nil
}
