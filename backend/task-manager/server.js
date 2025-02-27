require("dotenv").config();
const express = require("express");
const morgan = require("morgan");
const cors = require("cors");
const routes = require("./routes");
const connectDB = require("./config/db");

const app = express();
const PORT = process.env.PORT || 5000;

// Middleware
app.use(express.json());
app.use(morgan("dev"));
app.use(cors());
app.use("/api", routes);

app.get("/", (req, res) => {
  res.send("✅ Your server is running and connected to MongoDB!");
});

const startServer = async () => {
  try {
    await connectDB(); // Wait for MongoDB to connect before starting the server
    app.listen(PORT, () => {
      console.log(`🚀 Server is running on port ${PORT}`);
    });
  } catch (error) {
    console.error("❌ Failed to start the server due to DB connection error:", error);
    process.exit(1); // Exit process if MongoDB connection fails
  }
};

startServer();
