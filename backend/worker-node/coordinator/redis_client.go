package coordinator

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// RedisClient handles Redis connections and operations
type RedisClient struct {
	client *redis.Client
	logger *zap.Logger
}

// NewRedisClient creates and initializes a new Redis client
func NewRedisClient(redisAddr string, logger *zap.Logger) (*RedisClient, error) {
	// Create Redis client with options
	client := redis.NewClient(&redis.Options{
		Addr:         redisAddr,         // Redis server address
		Password:     "",                // No password by default
		DB:           0,                 // Use default DB
		DialTimeout:  5 * time.Second,   // Connection timeout
		ReadTimeout:  3 * time.Second,   // Read timeout
		WriteTimeout: 3 * time.Second,   // Write timeout
		PoolSize:     10,                // Connection pool size
	})
	
	// Verify connection works
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Ping Redis to check the connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	
	logger.Info("Connected to Redis", zap.String("address", redisAddr))
	
	return &RedisClient{
		client: client,
		logger: logger,
	}, nil
}

// Close cleanly shuts down the Redis connection
func (rc *RedisClient) Close() error {
	return rc.client.Close()
}

// GetClient returns the underlying Redis client
func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
}