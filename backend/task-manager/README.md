### **README for Node.js Task Manager Service**  

---

# **Node.js Task Manager Service (Task Orchestrator) 🚀**  

This service is the **backend task manager** for the **Decentralized Distributed Compute Network (DDCN)**. It handles **task submission, worker assignment, result verification, and reward distribution**.  

---

## **📌 Features**  

✅ Accepts tasks from users via REST API  
✅ Assigns tasks to available Go-based gRPC worker nodes  
✅ Stores task details & results in MongoDB  
✅ Verifies task completion & correctness  
✅ Handles worker node rewards  

---

## **📂 Folder Structure**  

```
/task-manager
│── /src
│   ├── /routes          # Express API routes
│   ├── /controllers     # Business logic for API endpoints
│   ├── /services        # gRPC client & MongoDB connection
│   ├── /models          # Task & Worker schemas (MongoDB)
│   ├── /utils           # Helper functions
│   ├── /config          # Configurations (DB, gRPC, env variables)
│   ├── /tests           # Unit & integration tests
│   ├── app.js           # Main Express application
│   ├── server.js        # Server initialization
│── /proto               # gRPC .proto definition file
│── .env                 # Environment variables
│── package.json         # Dependencies & scripts
│── README.md            # Project documentation
```

---

## **🛠 Prerequisites**  

📌 **Install Dependencies**  
```sh
npm install express mongoose grpc @grpc/proto-loader dotenv axios
```

📌 **MongoDB Setup** (Ensure it's running locally or in a cloud instance)  
```sh
mongod --dbpath ./data/db
```

---

## **🌍 REST API Endpoints (Express.js)**  

| Method | Endpoint               | Description |
|--------|------------------------|-------------|
| `POST` | `/tasks/submit`        | Accepts task requests from users |
| `GET`  | `/tasks/:task_id`      | Retrieves task status & results |
| `GET`  | `/workers/status`      | Gets available worker nodes |
| `POST` | `/tasks/verify`        | Verifies task correctness |
| `POST` | `/workers/reward`      | Distributes rewards to worker nodes |

📤 **Example Task Submission Request (JSON)**  
```json
{
  "task_id": "001",
  "type": "hashing",
  "input": {
    "algorithm": "SHA-256",
    "data": "Hello, DDCN!"
  }
}
```

---

## **🖥 gRPC Worker Communication (Go Workers)**  

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

📌 **Generate gRPC Code**  
```sh
npx grpc_tools_node_protoc --js_out=import_style=commonjs,binary:. \
  --grpc_out=. --proto_path=./proto ./proto/task.proto
```

📌 **Calling a Worker Node (Task Assignment Example)**  
```js
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const packageDefinition = protoLoader.loadSync("./proto/task.proto");
const taskProto = grpc.loadPackageDefinition(packageDefinition).Worker;

const client = new taskProto("localhost:50051", grpc.credentials.createInsecure());

client.ProcessTask({ task_id: "001", type: "hashing", input_data: "Hello, World!" }, (err, response) => {
    if (err) console.error(err);
    else console.log("Task Result:", response);
});
```

---

## **📦 Database (MongoDB Schema)**  

📌 **Task Schema (`/models/task.js`)**  
```js
const mongoose = require("mongoose");

const TaskSchema = new mongoose.Schema({
  task_id: { type: String, required: true, unique: true },
  type: { type: String, required: true },
  input: { type: Object, required: true },
  result: { type: String, default: null },
  status: { type: String, enum: ["pending", "processing", "completed"], default: "pending" },
});

module.exports = mongoose.model("Task", TaskSchema);
```

📌 **Worker Schema (`/models/worker.js`)**  
```js
const mongoose = require("mongoose");

const WorkerSchema = new mongoose.Schema({
  worker_id: { type: String, required: true, unique: true },
  status: { type: String, enum: ["idle", "busy"], default: "idle" },
  last_task: { type: String, default: null },
});

module.exports = mongoose.model("Worker", WorkerSchema);
```

---

## **🚀 Running the Task Manager**  

📌 **Start MongoDB**  
```sh
mongod --dbpath ./data/db
```

📌 **Run the Node.js Service**  
```sh
npm start
```

📌 **Test API with Postman or cURL**  
```sh
curl -X POST http://localhost:3000/tasks/submit -H "Content-Type: application/json" -d '{
  "task_id": "001",
  "type": "hashing",
  "input": { "algorithm": "SHA-256", "data": "Hello, DDCN!" }
}'
```

📌 **Worker gRPC Call Example**  
```sh
node workerClient.js
```
