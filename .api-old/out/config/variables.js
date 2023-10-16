"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.config = exports.PathType = void 0;
var PathType;
(function (PathType) {
    PathType["Local"] = "local";
    PathType["Nfs"] = "nfs";
    PathType["Cifs"] = "cifs";
})(PathType || (exports.PathType = PathType = {}));
exports.config = {
    PORT: parseInt(process.env.PORT) || 3000,
    DOCKER_API_URL: process.env.DOCKER_API_URL || "localhost:2375",
    CONTAINER_IMAGE: "alpine:latest",
    BACKUP_RETENTION: parseInt(process.env.BACKUP_RETENTION) || 7,
    BACKUP_PATH: process.env.BACKUP_PATH || "/backup",
    BACKUP_PATH_TYPE: process.env.BACKUP_PATH_TYPE || PathType.Local,
    BACKUP_PATH_USERNAME: process.env.BACKUP_PATH_USERNAME || "",
    BACKUP_PATH_PASSWORD: process.env.BACKUP_PATH_PASSWORD || "",
};
//# sourceMappingURL=variables.js.map