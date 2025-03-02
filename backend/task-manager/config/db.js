require("dotenv").config();
const mongoose = require("mongoose");


const connectDB = async (retries = 5, delay = 5000) => {

  const mongoURI = process.env.NODE_ENV === 'docker' 
    ? 'mongodb://mongodb:27017/taskDB'  // Docker environment
    : process.env.MONGO_URI;            // Local development (from .env)

  while (retries) {
    try {
      await mongoose.connect(mongoURI, {
        useNewUrlParser: true,
        useUnifiedTopology: true,
      });
      console.log("✅ MongoDB Connected");
      return;
    } catch (error) {
      console.error(`❌ MongoDB Connection Failed: ${error.message}`);
      retries -= 1;
      console.log(`🔄 Retrying in ${delay / 1000} seconds... (${retries} attempts left)`);
      await new Promise((res) => setTimeout(res, delay));
    }
  }

  console.error("❌ MongoDB connection failed after multiple attempts. Exiting...");
  process.exit(1);
};

module.exports = connectDB;
