package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"context"
	pb "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/hashing"
)

// HashingServiceServer implements the gRPC service
type HashingServiceServer struct {
	pb.UnimplementedHashingServiceServer
}

// ComputeSHA256 computes SHA-256 hash
func (s *HashingServiceServer) ComputeSHA256(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	hash := sha256.Sum256([]byte(req.Input))
	return &pb.HashResponse{Hash: hex.EncodeToString(hash[:])}, nil
}
