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
exports.init = void 0;
// Import libraries
const axios_1 = __importDefault(require("axios"));
//Import global variables
const variables_1 = require("../config/variables");
// ---------- External funcitons ----------
function init() {
    return __awaiter(this, void 0, void 0, function* () {
        console.log("Initializing program");
        // Pulling image for the backup container
        console.log(`Pulling image: ${variables_1.config.CONTAINER_IMAGE}`);
        const response = yield pullImage();
        if (response) {
            console.log(`Successfully pulld image: ${variables_1.config.CONTAINER_IMAGE}`);
        }
        else {
            throw new Error(`There was an error while pulling image: ${variables_1.config.CONTAINER_IMAGE}`);
        }
    });
}
exports.init = init;
// ---------- Internal functions ----------
// Pull the Docker image used for creating the Backup
function pullImage() {
    return __awaiter(this, void 0, void 0, function* () {
        const pullUrl = `http://${variables_1.config.DOCKER_API_URL}/images/create`;
        const params = {
            fromImage: `${variables_1.config.CONTAINER_IMAGE}`,
            //tag: 'latest',
        };
        // Send the API request to pull the image
        const response = yield axios_1.default.post(pullUrl, null, { params });
        //console.log(response)
        // Check the response for success or handle any errors
        return response.status === 200;
    });
}
//# sourceMappingURL=init.js.map