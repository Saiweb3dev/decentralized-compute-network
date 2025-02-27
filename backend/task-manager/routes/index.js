const express = require("express");
const router = express.Router();
const controller = require("../controllers/controllerEndpoint");

router.get("/test", controller.testController.test);

module.exports = router;
