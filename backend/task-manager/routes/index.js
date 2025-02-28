const express = require("express");
const router = express.Router();
const controller = require("../controllers/controllerEndpoint");

router.get("/test", controller.testController.test);

// Task routes
router.post("/task",controller.taskController.createTask);
router.get("/task/:taskId",controller.taskController.getTaskById);
router.get("/task/user/:userId",controller.taskController.getTasksByUserId);
router.put("/task/:taskId",controller.taskController.updateTask);
router.delete("/task/:taskId",controller.taskController.deleteTask);


module.exports = router;
