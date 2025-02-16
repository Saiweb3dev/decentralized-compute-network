 ## **⚙️ Supported Computation Types & Task Formats**  

### **1️⃣ Hashing (SHA-256, Keccak-256, Blake2b, etc.)**  
🔹 **Use Case**: Data integrity, cryptographic verification  
🔹 **Supported Algorithms**: `SHA-256`, `Keccak-256`, `Blake2b`  

📤 **Example Task Submission:**  
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
📥 **Worker Node Output:**  
```json
{
  "task_id": "001",
  "status": "completed",
  "result": "dffd6021bb2bd5b0af6767d4c54f70c2c86d146149f6b1b6c88cd0d6a18f5a99"
}
```

✅ **Verification**: Backend recomputes the hash and compares the results.  

---

### **2️⃣ Encryption/Decryption (AES-256, RSA-2048)**  
🔹 **Use Case**: Secure messaging, encrypted storage  
🔹 **Supported Algorithms**: `AES-256`, `RSA-2048`  

📤 **Example Task Submission:**  
```json
{
  "task_id": "002",
  "type": "encryption",
  "input": {
    "algorithm": "AES-256",
    "mode": "encrypt",
    "key": "securepassword123",
    "data": "Confidential Information"
  }
}
```
📥 **Worker Node Output:**  
```json
{
  "task_id": "002",
  "status": "completed",
  "result": "EncryptedString=="
}
```

✅ **Verification**: Backend decrypts the result using the provided key to check correctness.  

---

### **3️⃣ Modular Exponentiation (Mod Exp for cryptographic computations)**  
🔹 **Use Case**: RSA key generation, Diffie-Hellman, cryptographic proofs  
🔹 **Supported Operation**: `base^exponent mod modulus`  

📤 **Example Task Submission:**  
```json
{
  "task_id": "003",
  "type": "mod_exp",
  "input": {
    "base": "5",
    "exponent": "65537",
    "modulus": "99999989"
  }
}
```
📥 **Worker Node Output:**  
```json
{
  "task_id": "003",
  "status": "completed",
  "result": "82927465"
}
```

✅ **Verification**: Backend recomputes `5^65537 mod 99999989` and compares with the worker's result.  

---

## **🔗 API Endpoints**  

| Method | Endpoint               | Description |
|--------|------------------------|-------------|
| `POST` | `/tasks/submit`        | Submit a new computation task |
| `GET`  | `/tasks/:task_id`      | Check task status & results |
| `GET`  | `/workers/status`      | Get active worker node details |
| `POST` | `/tasks/verify`        | Verify worker task results |
| `POST` | `/workers/reward`      | Reward workers for completed tasks |
