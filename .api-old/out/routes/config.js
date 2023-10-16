"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// Import libraries
const express_1 = __importDefault(require("express"));
// Import controller functions
const config_1 = require("../controllers/config");
const router = express_1.default.Router();
// Return current config
router.get("/", (reg, res) => {
    res.json((0, config_1.getConfig)());
});
exports.default = router;
//# sourceMappingURL=config.js.map