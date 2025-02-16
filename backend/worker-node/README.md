### **README for Go Worker Service**  

---

# **Go Worker Service (Computation Node) 🚀**  

This service is a **decentralized worker node** in the **Decentralized Distributed Compute Network (DDCN)**. It receives computation tasks from the **Node.js Task Manager** via gRPC, processes them, and returns results.  

---

## **📌 Features**  

✅ Receives & processes tasks via gRPC  
✅ Supports multiple computation types (hashing, encryption, modular exponentiation)  
✅ Uses efficient CPU utilization techniques  
✅ Returns results to the Task Manager  
✅ Can run multiple worker nodes independently  

---

## **📂 Folder Structure**  

```
/worker-service
│── /proto               # gRPC .proto definition file
│── /config              # Configuration files (ports, env variables)
│── /handlers            # Task processing logic
│── /utils               # Helper functions for algorithms
│── /server.go           # gRPC server implementation
│── /main.go             # Worker node entry point
│── go.mod               # Go module dependencies
│── README.md            # Project documentation
```

---

## **🛠 Prerequisites**  

📌 **Install Go & Dependencies**  
```sh
sudo apt update && sudo apt install -y golang  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
```

📌 **Install Required Packages**  
```sh
go mod init worker-service  
go get google.golang.org/grpc  
go get google.golang.org/protobuf  
go get golang.org/x/crypto  
```

📌 **Ensure gRPC is installed**  
```sh
protoc --version  
```

---

## **📡 gRPC API for Worker Service**  

📌 **Proto Definition (`/proto/task.proto`)**  
```proto
syntax = "proto3";

service Worker {
  rpc ProcessTask (TaskRequest) returns (TaskResponse);
}

message TaskRequest {
  string task_id = 1;
  string type = 2;
  string input_data = 3;
}

message TaskResponse {
  string task_id = 1;
  string status = 2;
  string result = 3;
}
```

📌 **Generate gRPC Code for Go**  
```sh
protoc --go_out=. --go-grpc_out=. proto/task.proto  
```

---

## **💻 Computation Types Supported**  

| **Computation Type**  | **Input Format**  | **Example**  |
|----------------------|----------------|--------------|
| **SHA-256 Hashing**  | `"Hello, World!"`  | `"b94d27b..."`  |
| **Keccak-256 Hashing**  | `"Hello, DDCN!"`  | `"a8b5c3f..."`  |
| **AES Encryption**  | `{ "key": "secret123", "plaintext": "data" }`  | `"EncryptedString123..."`  |
| **Modular Exponentiation**  | `{ "base": 5, "exponent": 3, "mod": 13 }`  | `"8"`  |

---

## **🖥 Implementing the gRPC Worker**  

📌 **Start a gRPC Worker Node (`/main.go`)**  
```go
package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    pb "worker-service/proto"
)

type server struct {
    pb.UnimplementedWorkerServer
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterWorkerServer(grpcServer, &server{})

    log.Println("Worker Node started on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
```

📌 **Handle Task Processing (`/handlers/task.go`)**  
```go
func (s *server) ProcessTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
    var result string

    switch req.Type {
    case "hashing":
        result = sha256Hash(req.InputData)
    case "encryption":
        result = aesEncrypt(req.InputData, "default_key")
    case "modexp":
        result = modExp(req.InputData)
    default:
        return &pb.TaskResponse{TaskId: req.TaskId, Status: "error", Result: "Unsupported Task"}, nil
    }

    return &pb.TaskResponse{TaskId: req.TaskId, Status: "completed", Result: result}, nil
}
```

---

## **🛠 Running the Worker Node**  

📌 **Start the Worker Node**  
```sh
go run main.go
```

📌 **Check if Worker is Listening**  
```sh
netstat -tulnp | grep 50051
```

📌 **Test gRPC Worker (Using Node.js Client for Now)**  
```sh
node workerClient.js
```

---

## **📦 Scaling Worker Nodes**  

✅ Run multiple workers for load balancing  
✅ Assign tasks using a round-robin scheduler in Task Manager  
✅ Use Docker & Kubernetes for distributed deployment  

📌 **Run Multiple Worker Instances**  
```sh
go run main.go --port=50052  
go run main.go --port=50053  
```

📌 **Monitor CPU & Memory Usage**  
```sh
htop
```

---

## **🎯 Next Steps**  

- Implement **worker load monitoring**  
- Add **GPU acceleration for heavy computations**  
- Deploy using **Docker & Kubernetes**  
