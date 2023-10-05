// Import libraries
import express from "express";

// Import controller functions
import {
    getDockerVolumeNames
} from "../controllers/backup"


const router = express.Router();

// ---------- Backup information ----------

// Get a list of all Docker volumes
router.get("/", async (req, res) => {
    try {
        const volumeNames = await getDockerVolumeNames();
        console.log(`Request on ${req.originalUrl}, Response was ${JSON.stringify(volumeNames)}`);
        res.json(volumeNames);
    } catch (error) {
        console.error(`Request on ${req.originalUrl}, Error was ${error.message}`);
        res.status(500)
        res.send(error.message)
    }
});

// 
router.get("/:volumeName", (req, res) => {

});

// ---------- Backup creation ----------

//
router.post("/create", (req, res) => {

});

//
router.post("/create/:volumeName", (req, res) => {

});

// ---------- Backup purge ----------

//
router.post("/purge", (req, res) => {

});

//
router.post("/purge/:volumeName", (req, res) => {

});

// ---------- Backup deletion ----------

//
router.delete("/delete", (req, res) => {

});

export default router;