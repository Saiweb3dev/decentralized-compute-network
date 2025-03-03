package interceptors

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/coordinator"
)

// LoggingInterceptor provides a gRPC interceptor that performs custom logging
// and coordinates with the node coordinator to track task activity
func LoggingInterceptor(logger *zap.Logger, coordinator *coordinator.NodeCoordinator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract method and service information
		fullMethodName := info.FullMethod
		service, method := parseMethodName(fullMethodName)
		
		// Log the request
		logger.Debug("Received request", 
			zap.String("service", service),
			zap.String("method", method))
		
		// Record the start time
		startTime := time.Now()
		
		// Notify coordinator that a task is starting
		coordinator.TaskStarted()
		
		// Execute the RPC method
		resp, err := handler(ctx, req)
		
		// Notify coordinator that task is complete
		coordinator.TaskCompleted()
		
		// Calculate duration
		duration := time.Since(startTime)
		durationMs := float64(duration.Milliseconds())
		
		// Get the gRPC status code
		code := codes.OK
		if err != nil {
			code = status.Code(err)
		}
		
		// Create a colored log message
		// Method in cyan, time in yellow, status in green/red
		var statusColor string
		if code == codes.OK {
			statusColor = "\x1b[32m" // Green
		} else {
			statusColor = "\x1b[31m" // Red
		}
		
		logMsg := fmt.Sprintf("\x1b[36m%s.%s\x1b[0m completed in \x1b[33m%.2fms\x1b[0m with status %s%s\x1b[0m", 
			service, method, durationMs, statusColor, code.String())
		
		// Log with appropriate level based on status
		if code == codes.OK {
			logger.Info(logMsg)
		} else {
			logger.Error(logMsg, zap.Error(err))
		}
		
		return resp, err
	}
}

// parseMethodName extracts service and method from the full method string
// The format is "/package.Service/Method"
func parseMethodName(fullMethod string) (string, string) {
	if len(fullMethod) == 0 {
		return "unknown", "unknown"
	}
	
	// Skip the first slash
	servicePath := fullMethod[1:]
	
	// Find the last slash to separate service and method
	for i := len(servicePath) - 1; i >= 0; i-- {
		if servicePath[i] == '/' {
			return servicePath[:i], servicePath[i+1:]
		}
	}
	
	return servicePath, "unknown"
}