// Import libraries
import express from "express";

// Import util functions
import { printError } from "../utils/error";

// Import controller functions
import {
    getVolumes,
    getVolume,
    deleteVolume
} from "../controllers/backup"


const router = express.Router();

router.get("/test", async (req, res) => [
    deleteVolume("my_custom_volume")
]);

// ---------- Backup information ----------

// Get a list of all Docker volumes with their type (local, nfs, cifs)
// TODO: return backup status
router.get("/list", async (req, res) => {
    try {
        const volumeNames = await getVolumes();
        console.log(`Request on: ${req.originalUrl}, Response was: ${JSON.stringify(volumeNames)}`);
        res.json(volumeNames);
    } catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${printError(error)}`);
        res.status(500);
        res.send(error);
    }
});

// Get a single Docker volume by name
// TODO: return backup status
router.get("/list/:volumeName", async (req, res) => {
    const { volumeName } = req.params

    try {
        const volumeNames = await getVolume(volumeName);
        console.log(`Request on: ${req.originalUrl}, Response was: ${JSON.stringify(volumeNames)}`);
        res.json(volumeNames);
    } catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${printError(error)}`);
        res.status(500);
        res.send(error);
    }
});

// ---------- Backup creation ----------

// Create backup of all Docker volumes
router.post("/create", async (req, res) => {
    try {
        const { nfs, cifs } = req.body
        console.log(`Request on: ${req.originalUrl}`);
        res.status(200);
    } catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${printError(error)}`);
        res.status(500);
        res.send(error);
    }
});

//
router.post("/create/:volumeName", (req, res) => {
    try {
        const { nfs, cifs } = req.body
        console.log(`Request on: ${req.originalUrl}`);
        res.status(200);
    } catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${printError(error)}`);
        res.status(500);
        res.send(error);
    }
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