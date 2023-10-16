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
exports.deleteVolume = exports.getVolume = exports.getVolumes = void 0;
// Import libraries
const axios_1 = __importDefault(require("axios"));
//Import global variables
const variables_1 = require("../config/variables");
// ---------- External functions ----------
// Return a list of all Docker volumes
function getVolumes() {
    return __awaiter(this, void 0, void 0, function* () {
        const url = `http://${variables_1.config.DOCKER_API_URL}/system/df`;
        const volumesJson = (yield axios_1.default.get(url)).data;
        const volumes = {
            Volumes: volumesJson.Volumes.map((volume) => ({
                Name: volume.Name,
                Size: volume.UsageData.Size,
                Type: volume.Options ? volume.Options.type || "local" : "local",
                Labels: volume.Labels
            })),
        };
        return volumes;
    });
}
exports.getVolumes = getVolumes;
// Return a single Docker volume
function getVolume(volumeName) {
    return __awaiter(this, void 0, void 0, function* () {
        const volumes = yield getVolumes();
        const filteredVolumes = volumes.Volumes.filter((volume) => volume.Name === volumeName);
        return {
            Volumes: filteredVolumes,
        };
    });
}
exports.getVolume = getVolume;
// ---------- Internal functions ----------
// Run a Docker container to backup a certain volume
function backupVolume(dataVolumeName, backupVolumeName) {
    return __awaiter(this, void 0, void 0, function* () {
        const data = {
            "Image": `${variables_1.config.CONTAINER_IMAGE}`,
            "Cmd": ["ash", "-c", `tar cf /dest/${dataVolumeName}.backup.tar /data/.`],
            "HostConfig": {
                "Binds": [`${dataVolumeName}:/data`, `${backupVolumeName}:/dest`],
                "AutoRemove": true,
            },
        };
        // Send the API request to create the container
        const createUrl = `http://${variables_1.config.DOCKER_API_URL}/containers/create`;
        const createResponse = yield axios_1.default.post(createUrl, data);
        const containerId = createResponse.data.Id;
        // Start the container
        const startUrl = `http://${variables_1.config.DOCKER_API_URL}/containers/${containerId}/start`;
        yield axios_1.default.post(startUrl);
        // The container will run and create the backup as specified in the command.
        // Optionally, you can wait for the container to finish if needed.
        // Wait for the container to finish (blocking)
        const waitUrl = `http://${variables_1.config.DOCKER_API_URL}/containers/${containerId}/wait`;
        yield axios_1.default.post(waitUrl);
    });
}
function createVolume(volumeName, type, path, username, password) {
    return __awaiter(this, void 0, void 0, function* () {
        let data;
        switch (type) {
            case variables_1.PathType.Local:
                data = {
                    "Name": `${volumeName}`,
                    "Driver": "local",
                    "DriverOpts": {
                        "type": "none",
                        "o": "bind",
                        "device": `${path}`
                    },
                    "Labels": {
                        "com.dvb.volume": true
                    }
                };
                break;
            case variables_1.PathType.Nfs:
                data = {};
                break;
            case variables_1.PathType.Cifs:
                data = {};
                break;
            default:
                data = {};
                break;
        }
        const response = yield axios_1.default.post(`http://${variables_1.config.DOCKER_API_URL}/volumes/create`, data);
        if (response.status !== 201) {
            throw new Error(response.data);
        }
    });
}
function deleteVolume(volumeName) {
    return __awaiter(this, void 0, void 0, function* () {
        const volumeToDelete = (yield axios_1.default.get(`http://${variables_1.config.DOCKER_API_URL}/volumes`)).data.Volumes.filter((volume) => volume.Name === volumeName);
        if (volumeToDelete[0].Labels["com.dvb.volume"]) {
            console.log("laft");
        }
    });
}
exports.deleteVolume = deleteVolume;
//# sourceMappingURL=backup.js.map