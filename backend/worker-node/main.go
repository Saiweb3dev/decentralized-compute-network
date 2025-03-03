package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/config"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/coordinator"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/interceptors"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/logger"
	"go.uber.org/zap"

	pb_hashing "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/hashing"
	pb_encryption "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/encryption"
	pb_modexp "github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/proto/modexp"

	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/hashing"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/encryption"
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/services/modexp"
)

func main() {
	// 1. Load application configuration
	cfg := config.LoadConfig()
	
	// 2. Initialize logger
	zapLogger, err := logger.InitLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer zapLogger.Sync()
	
	zapLogger.Info("Starting worker node", 
		zap.String("nodeId", cfg.NodeID),
		zap.String("address", cfg.NodeAddress))
	
	// 3. Initialize Redis client
	redisClient, err := coordinator.NewRedisClient(cfg.RedisAddress, zapLogger)
	if err != nil {
		zapLogger.Fatal("Failed to connect to Redis", zap.Error(err))
	}
	defer redisClient.Close()
	
	// 4. Initialize node coordinator
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	nodeCoordinator := coordinator.NewNodeCoordinator(
		cfg.NodeID,
		cfg.NodeAddress,
		redisClient,
		cfg.MaxTasks,
		zapLogger,
	)
	
	// 5. Start the coordinator
	nodeCoordinator.Start(ctx)
	
	// 6. Create gRPC server with our custom logging interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.LoggingInterceptor(zapLogger, nodeCoordinator)),
	)
	
	// 7. Register services
	pb_hashing.RegisterHashingServiceServer(grpcServer, &hashing.HashingServiceServer{})
	pb_encryption.RegisterEncryptionServiceServer(grpcServer, &encryption.EncryptionServiceServer{})
	pb_modexp.RegisterModExpServiceServer(grpcServer, &modexp.ModExpServiceServer{})
	
	// 8. Enable gRPC reflection for debugging
	reflection.Register(grpcServer)
	
	// 9. Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// 10. Start the server
	lis, err := net.Listen("tcp", cfg.NodeAddress)
	if err != nil {
		zapLogger.Fatal("Failed to listen", zap.Error(err))
	}
	
	// 11. Start server in a goroutine so the signal handling doesn't block
	go func() {
		zapLogger.Info(fmt.Sprintf("🚀 gRPC Worker Node %s running on %s", cfg.NodeID, cfg.NodeAddress))
		if err := grpcServer.Serve(lis); err != nil {
			zapLogger.Fatal("Failed to serve", zap.Error(err))
		}
	}()
	
	// 12. Wait for termination signal
	sig := <-sigChan
	zapLogger.Info("Received termination signal", zap.String("signal", sig.String()))
	
	// 13. Graceful shutdown
	zapLogger.Info("Shutting down gracefully...")
	nodeCoordinator.Stop()
	grpcServer.GracefulStop()
	zapLogger.Info("Worker node shutdown complete")
}