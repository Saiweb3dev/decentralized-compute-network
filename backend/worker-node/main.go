package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb_hashing "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/hashing"
	pb_encryption "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/encryption"
	pb_modexp "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/modexp"

	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/hashing"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/encryption"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/modexp"
)

// CustomLoggerInterceptor is a UnaryServerInterceptor that provides simplified logging
func CustomLoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		
		// Extract method and service from the full method name (format: "/service.name/method")
		fullMethodName := info.FullMethod
		var service, method string
		
		// Simple parsing of the full method name
		for i := 1; i < len(fullMethodName); i++ {
			if fullMethodName[i] == '/' {
				service = fullMethodName[1:i]
				method = fullMethodName[i+1:]
				break
			}
		}
		
		// Process the request
		resp, err := handler(ctx, req)
		
		// Calculate duration
		duration := time.Since(startTime)
		durationMs := float64(duration.Milliseconds())
		
		// Create a simplified and colorful log message
		status := "OK"
		if err != nil {
			status = "ERROR"
		}
		
		// Log with colors and only the fields we care about
		logger.Info(fmt.Sprintf("\x1b[36m%s.%s\x1b[0m completed in \x1b[33m%.2fms\x1b[0m with status \x1b[32m%s\x1b[0m", 
			service, method, durationMs, status))
		
		return resp, err
	}
}

func main() {
	// Custom zap logger config
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to build logger: %v", err)
	}
	defer logger.Sync()
	
	// Create gRPC server with our custom logging interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(CustomLoggerInterceptor(logger)),
	)

	// Register services
	pb_hashing.RegisterHashingServiceServer(grpcServer, &hashing.HashingServiceServer{})
	pb_encryption.RegisterEncryptionServiceServer(grpcServer, &encryption.EncryptionServiceServer{})
	pb_modexp.RegisterModExpServiceServer(grpcServer, &modexp.ModExpServiceServer{})

	// Enable gRPC reflection for debugging
	reflection.Register(grpcServer)

	// Start server
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("🚀 gRPC Server running on port 5001...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}