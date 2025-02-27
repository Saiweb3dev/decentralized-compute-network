const express = require("express");
const router = express.Router();
const controller = require("../controllers/controllerEndpoint");

router.get("/test", controller.testController.test);
router.post("/tasks",controller.taskController.createTask);

module.exports = router;
