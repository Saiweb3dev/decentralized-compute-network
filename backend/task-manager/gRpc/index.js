const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');
const loadProto = (protoFile) => {
    const packageDefinition = protoLoader.loadSync("" + path.join("../", "../backend/proto/", protoFile), {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
    return grpc.loadPackageDefinition(packageDefinition);
};

// Load all proto services
const hashingProto = loadProto('hashing.proto').hashing;
const encryptionProto = loadProto('encryption.proto').encryption;

// Create gRPC clients
const hashingClient = new hashingProto.HashingService(
    'localhost:5001',
    grpc.credentials.createInsecure()
);

const encryptionClient = new encryptionProto.EncryptionService(
    'localhost:5001',
    grpc.credentials.createInsecure()
);

module.exports = { hashingClient, encryptionClient };
