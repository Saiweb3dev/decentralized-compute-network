require("dotenv").config();
const mongoose = require("mongoose");


const connectDB = async (retries = 5, delay = 5000) => {
  while (retries) {
    try {
      await mongoose.connect(process.env.MONGO_URI, {
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
