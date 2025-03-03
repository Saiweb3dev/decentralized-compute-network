package coordinator

import (
	"time"
)

// NodeStatus represents the current state of a worker node
// This information is shared across the network to facilitate task distribution
type NodeStatus struct {
	// NodeID uniquely identifies this worker node
	NodeID string `json:"nodeId"`
	
	// Address is the gRPC endpoint where this node can be reached
	Address string `json:"address"`
	
	// CPULoad is the current CPU utilization percentage (0-100)
	CPULoad float64 `json:"cpuLoad"`
	
	// MemoryUsage is the current memory utilization percentage (0-100)
	MemoryUsage float64 `json:"memoryUsage"`
	
	// TasksRunning is the number of tasks currently being processed
	TasksRunning int `json:"tasksRunning"`
	
	// MaxTasks is the maximum number of concurrent tasks this node can handle
	MaxTasks int `json:"maxTasks"`
	
	// LastUpdated is the timestamp when this status was last refreshed
	LastUpdated time.Time `json:"lastUpdated"`
}