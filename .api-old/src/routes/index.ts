import express from "express";

import backupRouter from "./backup";
import configRouter from "./config";

const router = express.Router();

router.use("/backup", backupRouter);
router.use("/config", configRouter);

export default router;