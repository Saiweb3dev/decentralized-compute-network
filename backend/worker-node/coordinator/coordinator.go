package coordinator

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/saiweb3dev/decentralized-compute-network/backend/worker-node/utils"
)

// NodeCoordinator handles worker node status management and communication
type NodeCoordinator struct {
	// Core identity and connection info
	nodeID        string
	address       string
	redisClient   *RedisClient
	
	// Task management
	maxTasks      int
	taskCounter   int
	statusMutex   sync.Mutex
	
	// Channels for coordination
	updateChannel chan struct{}
	stopChannel   chan struct{}
	
	// Logging
	logger        *zap.Logger
}

// NewNodeCoordinator creates a new coordinator for this worker node
func NewNodeCoordinator(nodeID string, address string, redisClient *RedisClient, maxTasks int, logger *zap.Logger) *NodeCoordinator {
	return &NodeCoordinator{
		nodeID:        nodeID,
		address:       address,
		redisClient:   redisClient,
		maxTasks:      maxTasks,
		taskCounter:   0,
		logger:        logger,
		updateChannel: make(chan struct{}, 10),
		stopChannel:   make(chan struct{}),
	}
}

// Start begins the coordination operations
func (nc *NodeCoordinator) Start(ctx context.Context) {
	// Register this node initially
	nc.publishStatus()
	
	// Start subscription to status updates from other nodes
	go nc.subscribeToUpdates(ctx)
	
	// Start periodic status updater (heartbeat)
	go nc.periodicStatusUpdater(ctx)
	
	// Start monitoring for update triggers
	go nc.handleUpdateTriggers(ctx)
	
	nc.logger.Info("Node coordinator started",
		zap.String("nodeId", nc.nodeID),
		zap.String("address", nc.address))
}

// Stop gracefully shuts down the coordinator
func (nc *NodeCoordinator) Stop() {
	close(nc.stopChannel)
	// Final status update to show we're going offline
	nc.statusMutex.Lock()
	defer nc.statusMutex.Unlock()
	// Mark node as inactive by setting max tasks to 0
	nc.maxTasks = 0
	nc.publishStatus()
}

// TaskStarted should be called when a new task begins
func (nc *NodeCoordinator) TaskStarted() {
	nc.statusMutex.Lock()
	defer nc.statusMutex.Unlock()
	
	nc.taskCounter++
	// Signal that status needs updating
	select {
	case nc.updateChannel <- struct{}{}:
		// Successfully sent update signal
	default:
		// Channel buffer is full, but we'll have periodic updates anyway
	}
}

// TaskCompleted should be called when a task finishes
func (nc *NodeCoordinator) TaskCompleted() {
	nc.statusMutex.Lock()
	defer nc.statusMutex.Unlock()
	
	if nc.taskCounter > 0 {
		nc.taskCounter--
	}
	// Signal that status needs updating
	select {
	case nc.updateChannel <- struct{}{}:
		// Successfully sent update signal
	default:
		// Channel buffer is full, but we'll have periodic updates anyway
	}
}

// GetNodeStatus collects the current status of this node
func (nc *NodeCoordinator) GetNodeStatus() NodeStatus {
	nc.statusMutex.Lock()
	defer nc.statusMutex.Unlock()
	
	// Collect real system metrics
	cpuLoad, memoryUsage := utils.GetSystemMetrics()
	
	return NodeStatus{
		NodeID:       nc.nodeID,
		Address:      nc.address,
		CPULoad:      cpuLoad,
		MemoryUsage:  memoryUsage,
		TasksRunning: nc.taskCounter,
		MaxTasks:     nc.maxTasks,
		LastUpdated:  time.Now(),
	}
}

// publishStatus publishes the node's current status to Redis
func (nc *NodeCoordinator) publishStatus() {
	status := nc.GetNodeStatus()
	
	statusJson, err := json.Marshal(status)
	if err != nil {
		nc.logger.Error("Failed to marshal node status", zap.Error(err))
		return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Publish to a channel for real-time updates
	err = nc.redisClient.GetClient().Publish(ctx, "worker_node_status", statusJson).Err()
	if err != nil {
		nc.logger.Error("Failed to publish status update", zap.Error(err))
		return
	}
	
	// Also store in a hash for persistence and discovery
	err = nc.redisClient.GetClient().HSet(ctx, "worker_nodes", nc.nodeID, statusJson).Err()
	if err != nil {
		nc.logger.Error("Failed to store node status", zap.Error(err))
		return
	}
	
	nc.logger.Debug("Published node status",
		zap.Int("tasksRunning", status.TasksRunning),
		zap.Float64("cpuLoad", status.CPULoad))
}

// subscribeToUpdates subscribes to other nodes' status updates
func (nc *NodeCoordinator) subscribeToUpdates(ctx context.Context) {
	pubsub := nc.redisClient.GetClient().Subscribe(ctx, "worker_node_status")
	defer pubsub.Close()
	
	nc.logger.Info("Subscribed to worker node updates")
	
	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			// Process update from another node
			var status NodeStatus
			if err := json.Unmarshal([]byte(msg.Payload), &status); err != nil {
				nc.logger.Error("Failed to unmarshal node status", zap.Error(err))
				continue
			}
			
			// Skip our own updates
			if status.NodeID == nc.nodeID {
				continue
			}
			
			nc.logger.Debug("Received status update from node", 
				zap.String("nodeId", status.NodeID),
				zap.Int("tasksRunning", status.TasksRunning),
				zap.Float64("cpuLoad", status.CPULoad))
				
		case <-ctx.Done():
			nc.logger.Info("Stopping status subscription")
			return
			
		case <-nc.stopChannel:
			nc.logger.Info("Stopping status subscription")
			return
		}
	}
}

// periodicStatusUpdater publishes status updates at regular intervals
func (nc *NodeCoordinator) periodicStatusUpdater(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			nc.publishStatus()
			
		case <-ctx.Done():
			return
			
		case <-nc.stopChannel:
			return
		}
	}
}

// handleUpdateTriggers processes update requests from task start/complete events
func (nc *NodeCoordinator) handleUpdateTriggers(ctx context.Context) {
	// Use a ticker to limit update frequency to avoid flooding
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	
	var pendingUpdate bool
	
	for {
		select {
		case <-nc.updateChannel:
			pendingUpdate = true
			
		case <-ticker.C:
			if pendingUpdate {
				nc.publishStatus()
				pendingUpdate = false
			}
			
		case <-ctx.Done():
			return
			
		case <-nc.stopChannel:
			return
		}
	}
}