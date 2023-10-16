"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// Import libraries
const express_1 = __importDefault(require("express"));
// Import util functions
const error_1 = require("../utils/error");
// Import controller functions
const backup_1 = require("../controllers/backup");
const router = express_1.default.Router();
router.get("/test", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    return [
        (0, backup_1.deleteVolume)("my_custom_volume")
    ];
}));
// ---------- Backup information ----------
// Get a list of all Docker volumes with their type (local, nfs, cifs)
// TODO: return backup status
router.get("/list", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const volumeNames = yield (0, backup_1.getVolumes)();
        console.log(`Request on: ${req.originalUrl}, Response was: ${JSON.stringify(volumeNames)}`);
        res.json(volumeNames);
    }
    catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${(0, error_1.printError)(error)}`);
        res.status(500);
        res.send(error);
    }
}));
// Get a single Docker volume by name
// TODO: return backup status
router.get("/list/:volumeName", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const { volumeName } = req.params;
    try {
        const volumeNames = yield (0, backup_1.getVolume)(volumeName);
        console.log(`Request on: ${req.originalUrl}, Response was: ${JSON.stringify(volumeNames)}`);
        res.json(volumeNames);
    }
    catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${(0, error_1.printError)(error)}`);
        res.status(500);
        res.send(error);
    }
}));
// ---------- Backup creation ----------
// Create backup of all Docker volumes
router.post("/create", (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const { nfs, cifs } = req.body;
        console.log(`Request on: ${req.originalUrl}`);
        res.status(200);
    }
    catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${(0, error_1.printError)(error)}`);
        res.status(500);
        res.send(error);
    }
}));
//
router.post("/create/:volumeName", (req, res) => {
    try {
        const { nfs, cifs } = req.body;
        console.log(`Request on: ${req.originalUrl}`);
        res.status(200);
    }
    catch (error) {
        console.error(`Request on: ${req.originalUrl}, Error was: ${(0, error_1.printError)(error)}`);
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
exports.default = router;
//# sourceMappingURL=backup.js.map