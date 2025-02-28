const Task = require("../models/Task");
const { v4: uuidv4 } = require("uuid");

// @route POST /api/task
const createTask = async (req, res) => {
  try {
    const { userId, type, input } = req.body;
    if (!userId || !type || !input) {
      return res.status(400).json({ message: "Please provide all the fields" });
    }

    //Generate a unique task_id
    const taskId = uuidv4();

    //Create a new task
    const task = new Task({
      task_id: taskId,
      user_id: userId,
      type,
      input,
      status: "pending",
    });
    await task.save();

    //Return taskId to the user
    res.status(201).json({ message: "Task created successfully", taskId });
  } catch (error) {
    console.error(error);
    res.status(500).json({
      message: "An error occurred while creating the task",
    });
  }
};

//@route GET /api/task/:taskId
const getTaskById = async (req, res) => {
  try {
    const { taskId } = req.params;
    const task = await Task.findOne({ task_id: taskId });
    if (!task) {
      return res.status(404).json({ message: "Task not found" });
    }
    res.status(200).json({ task });
  } catch (error) {
    console.error(error);
    res.status(500).json({
      message: "An error occurred while retrieving the task",
    });
  }
};

//@route GET /api/task/user/:userId
const getTasksByUserId = async (req, res) => {
  try {
    const { userId } = req.params;
    const tasks = await Task.find({ user_id: userId });
    res.status(200).json({ tasks });
  } catch (err) {
    console.error("Error fetching user tasks:", err);
    res.status(500).json({
      message: "An error occurred while retrieving the tasks",
    });
  }
};

//@route PUT /api/tasks/:taskId
const updateTask = async (req, res) => {
  try {
    const { taskId } = req.params;
    const { status, result } = req.body;

    if (!status) {
      return res.status(400).json({ message: "Status is required" });
    }
    const task = await Task.findOneAndUpdate(
      { task_id: taskId },
      { status: status, result: result },
      { new: true }
    );
    if (!task) {
      return res.status(404).json({ message: "Task not found" });
    }
    res.status(200).json({ message: "Task updated successfully", task });
  } catch (err) {
    console.error("Error updating task:", err);
    res.status(500).json({
      message: "An error occurred while updating the task",
    });
  }
};

// @route DELETE /api/tasks/:taskId
const deleteTask = async (req, res) => {
  try {
    const { taskId } = req.params;

    //Check if task exists
    const task = await Task.findOne({ task_id: taskId });
    if (!task) {
      return res.status(404).json({ message: "Task not found" });
    }

    if (task.status !== "pending") {
      return res
        .status(400)
        .json({ message: "Cannot delete a processing or completed task" });
    }

    await Task.deleteOne({ task_id: taskId });
    res.status(200).json({ message: "Task deleted successfully" });
  } catch (err) {
    console.error("Error deleting task:", err);
    res.status(500).json({
      message: "An error occurred while deleting the task",
    });
  }
};

module.exports = { 
  createTask,
  getTaskById,
  getTasksByUserId,
  updateTask,
  deleteTask,
 };
