const Task = require("../models/Task");

// POST /api/tasks - Create a new task
const createTask = async (req, res) => {
  try {
    const { title, description } = req.body;
    if (!title || !description) {
      return res.status(400).json({ message: "Title and description are required" });
    }

    const task = new Task({ title, description });
    await task.save();

    res.status(201).json({ message: "Task created successfully", task });
  } catch (error) {
    console.error("Error creating task:", error);
    res.status(500).json({ message: "Internal Server Error" });
  }
};

module.exports = { createTask };
