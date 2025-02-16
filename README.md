# **📌 Decentralized Distributed Compute Network (DDCN)**  
> **A gRPC & Docker-powered decentralized system for distributed task execution with Ethereum smart contracts for rewards.**  

---

## **🎯 Project Overview**  
The **Decentralized Distributed Compute Network (DDCN)** allows users to submit computational tasks (e.g., solving a math puzzle). Decentralized worker nodes process these tasks, verify results, and store them in IPFS. A **smart contract** ensures that workers receive fair rewards for successful execution.  

### **📌 How It Works?**
1. **User submits a task** via the frontend UI (Next.js).  
2. The **Task Manager (Node.js REST API)** stores the task in **MongoDB** and assigns it to an available **worker node (gRPC in Go)**.  
3. The **worker node processes the task** and submits the result back to the **Task Manager**.  
4. The result is stored in **IPFS**, and the hash is recorded in **MongoDB**.  
5. The **Ethereum smart contract verifies execution** and **rewards the worker** upon successful completion.  

---

## **📁 Project Structure**
```
Decentralized_Distributed_Compute_Network/
│── frontend/                 # Next.js frontend for user interaction
│── backend/                  # Backend microservices
│   ├── task-manager/         # Node.js REST API to manage tasks
│   ├── worker-node/          # Go gRPC server for distributed execution
│   ├── smart-contracts/      # Solidity smart contracts for reward distribution
│   ├── db/                   # MongoDB database
│── docker-compose.yml        # Docker script to set up services
│── README.md                 # This file
```

---

## **🛠 Prerequisite Tools**
Before running the project, ensure the following are installed:
- **[Node.js](https://nodejs.org/)**
- **[Go](https://go.dev/)**
- **[Docker](https://www.docker.com/)**
- **[MongoDB](https://www.mongodb.com/)**
- **[IPFS](https://ipfs.tech/)**
- **[Hardhat](https://hardhat.org/)** for smart contract development

---

## **🔹 Microservices Overview**
| Service           | Technology  | Description |
|------------------|------------|-------------|
| **Frontend**      | Next.js    | User interface to submit tasks |
| **Task Manager**  | Node.js, Express, MongoDB | Handles task submissions and worker coordination |
| **Worker Node**   | Go, gRPC   | Decentralized worker nodes that process tasks |
| **IPFS Storage**  | IPFS       | Stores processed task results |
| **Smart Contracts** | Solidity, Hardhat | Manages reward distribution for workers |

---

## **📌 API Endpoints (Task Manager - Node.js)**
| Method | Endpoint                  | Description |
|--------|---------------------------|-------------|
| **POST**  | `/submit-task`         | User submits a task to be processed |
| **GET**   | `/tasks`               | Fetch all submitted tasks |
| **POST**  | `/assign-task`         | Assigns a task to an available worker |
| **POST**  | `/submit-result`       | Worker submits task execution result |
| **GET**   | `/task-result/:id`     | Retrieve task result from IPFS |

---

## **🖧 gRPC Endpoints (Worker Node - Go)**
| Method        | Service Name         | Description |
|--------------|----------------------|-------------|
| `ProcessTask` | `WorkerService`      | Processes a given task and returns result |
| `GetStatus`   | `WorkerService`      | Returns worker node availability status |

---

## **📌 Running the Project with Docker**
To easily deploy all services, use Docker.

### *1️⃣ Install Docker & Docker Compose**
```sh
# Install Docker (Linux)
sudo apt install -y docker.io docker-compose
```

### **2️⃣ Clone the Repository**
```sh
git clone https://github.com/your-username/dten-math-solver.git
cd dten-math-solver
```

### **3️⃣ Start the System**
```sh
docker-compose up --build
```

### **4️⃣ Access the Services**
- **Frontend UI**: `http://localhost:3000`
- **Task Manager API**: `http://localhost:5000`
- **MongoDB**: `mongodb://localhost:27017`
- **IPFS Web UI**: `http://localhost:5001/webui`
- **Hardhat Local Node**: `http://localhost:8545`

## **🔹 Reward Distribution via Smart Contract**
Once a worker completes a task:
1. The **task hash is verified** on-chain.
2. If verified, the **Ethereum smart contract** distributes rewards.
3. The worker receives **ETH** or tokens as a reward.

**Smart Contract Deployment (sepolia Testnet)**
```sh
npx hardhat run --network sepolia scripts/deploy.js
```

---

## **🎯 Future Enhancements**
✅ **Multi-worker support for load balancing**  
✅ **On-chain governance for task validation**  
✅ **Additional computational task types (AI, data processing)**  
