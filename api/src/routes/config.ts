/*
*   Config router
*
*
*/

// Import libraries
import express from "express";

// Import controller functions
import {
    getConfig
} from "../controllers/config"


const router = express.Router();


// Return current config
router.get("/", (reg, res) => {
    res.json(getConfig());
});

export default router;