const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

// Load the .proto file
const packageDefinition = protoLoader.loadSync("../proto/hello.proto", {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const helloProto = grpc.loadPackageDefinition(packageDefinition).hello;

// Create a gRPC client
const client = new helloProto.HelloService(
  "localhost:5001",
  grpc.credentials.createInsecure()
);

// Send a request every 1 second
setInterval(() => {
  client.SayHelloBro({ message: "Hello from Express!" }, (error, response) => {
    if (error) {
      console.error("Error:", error);
    } else {
      console.log("Server Response:", response.reply);
    }
  });
}, 1000);
