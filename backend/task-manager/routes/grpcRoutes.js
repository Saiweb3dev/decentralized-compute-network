const express = require("express");
const { hashingClient, encryptionClient } = require("../gRpc");

const router = express.Router();

router.get("/test", (req, res) => {
  res.json({ message: "GRPC is working!" });
});

// 🔹 Compute SHA-256 Hash
router.post("/hashing/sha256", (req, res) => {
    const { input } = req.body;

    hashingClient.ComputeSHA256({ input }, (error, response) => {
        if (error) {
            console.error("gRPC Error:", error);
            return res.status(500).json({ error: "gRPC Error" });
        }
        res.json({ hash: response.hash });
    });
});

// 🔹 Encrypt using AES-256
router.post("/encryption/aes256/encrypt", (req, res) => {
    const { plaintext, key } = req.body;

    encryptionClient.EncryptAES256({ plaintext, key }, (error, response) => {
        if (error) {
            console.error("gRPC Error:", error);
            return res.status(500).json({ error: "gRPC Error" });
        }
        res.json({ ciphertext: response.ciphertext });
    });
});

// 🔹 Decrypt using AES-256
router.post("/encryption/aes256/decrypt", (req, res) => {
    const { ciphertext, key } = req.body;

    encryptionClient.DecryptAES256({ ciphertext, key }, (error, response) => {
        if (error) {
            console.error("gRPC Error:", error);
            return res.status(500).json({ error: "gRPC Error" });
        }
        res.json({ decrypted: response.plaintext });
    });
});

module.exports = router;
