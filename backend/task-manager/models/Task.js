const mongoose = require("mongoose");

const TaskSchema = new mongoose.Schema({
  task_id: { type: String, required: true, unique: true },
  type: { type: String, required: true },
  input: { type: Object, required: true },
  result: { type: String, default: null },
  status: { type: String, enum: ["pending", "processing", "completed"], default: "pending" },
});

module.exports = mongoose.model("Task", TaskSchema);
