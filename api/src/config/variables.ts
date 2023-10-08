interface Configuration {
    PORT: number;
    DOCKER_API_URL: string;
    readonly CONTAINER_IMAGE: string;
    BACKUP_RETENTION: number;
    BACKUP_PATH: string;
    BACKUP_PATH_TYPE: PathType;
    BACKUP_PATH_USERNAME: string
    BACKUP_PATH_PASSWORD: string
}

enum PathType {
    Local = "local",
    Nfs = "nfs",
    Cifs = "cifs",
}

export const config: Configuration = {
    PORT: parseInt(process.env.PORT) || 3000,
    DOCKER_API_URL: process.env.DOCKER_API_URL || "localhost:2375",
    CONTAINER_IMAGE: "alpine:latest",
    BACKUP_RETENTION: parseInt(process.env.BACKUP_RETENTION) || 7,
    BACKUP_PATH: process.env.BACKUP_PATH|| "/backup",
    BACKUP_PATH_TYPE: process.env.BACKUP_PATH_TYPE as PathType || PathType.Local,
    BACKUP_PATH_USERNAME: process.env.BACKUP_PATH_USERNAME || "",
    BACKUP_PATH_PASSWORD: process.env.BACKUP_PATH_PASSWORD || "",
}