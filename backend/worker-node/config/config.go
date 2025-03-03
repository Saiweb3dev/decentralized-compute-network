package config

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)

// Config holds all configuration for the worker node
type Config struct {
	// NodeID uniquely identifies this worker node in the network
	NodeID string
	
	// NodeAddress is the gRPC server address in host:port format
	NodeAddress string
	
	// RedisAddress is the Redis server address in host:port format
	RedisAddress string
	
	// MaxTasks is the maximum number of concurrent tasks this node can handle
	MaxTasks int
	
	// LogLevel defines the minimum log level to display (debug, info, warn, error)
	LogLevel string
}

// LoadConfig loads configuration from environment variables with sensible defaults
func LoadConfig() *Config {
	// Seed random number generator for ID generation
	rand.Seed(time.Now().UnixNano())
	
	// Load NodeID from environment or generate a random one
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID = fmt.Sprintf("worker-%d", rand.Intn(10000))
	}
	
	// Load NodeAddress from environment or use default
	nodeAddress := os.Getenv("NODE_ADDRESS")
	if nodeAddress == "" {
		nodeAddress = "0.0.0.0:5001" // Default to listen on all interfaces
	}
	
	// Load Redis address from environment or use default
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379" // Default Redis port
	}
	
	// Load max tasks setting from environment or use default
	maxTasksStr := os.Getenv("MAX_TASKS")
	maxTasks := 10 // Default value
	if maxTasksStr != "" {
		if val, err := strconv.Atoi(maxTasksStr); err == nil && val > 0 {
			maxTasks = val
		}
	}
	
	// Load log level from environment or use default
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info" // Default log level
	}
	
	return &Config{
		NodeID:       nodeID,
		NodeAddress:  nodeAddress,
		RedisAddress: redisAddress,
		MaxTasks:     maxTasks,
		LogLevel:     logLevel,
	}
}