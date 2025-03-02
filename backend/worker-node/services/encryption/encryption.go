package encryption

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	pb "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/encryption"
)

// EncryptAES256 encrypts data using AES-256
func EncryptAES256(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // AES-GCM nonce
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// EncryptionServiceServer struct
type EncryptionServiceServer struct {
	pb.UnimplementedEncryptionServiceServer
}

// EncryptAES256 gRPC method
func (s *EncryptionServiceServer) EncryptAES256(ctx context.Context, req *pb.EncryptionRequest) (*pb.EncryptionResponse, error) {
	encrypted, err := EncryptAES256(req.Key, req.Plaintext)
	if err != nil {
		return nil, errors.New("encryption failed")
	}
	return &pb.EncryptionResponse{Ciphertext: encrypted}, nil
}
