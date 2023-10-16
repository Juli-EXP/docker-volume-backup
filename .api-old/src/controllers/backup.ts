/*
*    Backup controller
*    Functions starting with docker are responsible for making API calls with the Docker socket
*
*/

// Import libraries
import axios from "axios";

//Import global variables
import { config, PathType } from "../config/variables";


// -------------------- External functions --------------------

// Return a list of all Docker volumes
export async function getVolumes() {
    const url = `http://${config.DOCKER_API_URL}/system/df`;
    const volumesJson = (await axios.get(url)).data;

    const volumes = {
        Volumes: volumesJson.Volumes.map((volume: any) => ({
            Name: volume.Name,
            Size: volume.UsageData.Size,
            Type: volume.Options ? volume.Options.type || "local" : "local",    // Checks if volume.Options and volume.Options.type exist. If not it sets the type to local
            Labels: volume.Labels
        })),
    };

    return volumes;
}

// Return a single Docker volume
export async function getVolume(volumeName: string) {
    const volumes = await getVolumes();

    const filteredVolumes = volumes.Volumes.filter((volume: any) => volume.Name === volumeName);

    return {
        Volumes: filteredVolumes,
    };
}


// -------------------- Docker API calls --------------------




// -------------------- Internal functions --------------------

// Run a Docker container to backup a certain volume
async function backupVolume(dataVolumeName: string, backupVolumeName: string) {
    const data = {
        "Image": `${config.CONTAINER_IMAGE}`,
        "Cmd": ["ash", "-c", `tar cf /dest/${dataVolumeName}.backup.tar /data/.`],
        "HostConfig": {
            "Binds": [`${dataVolumeName}:/data`, `${backupVolumeName}:/dest`],
            "AutoRemove": true,
        },
    };

    // Send the API request to create the container
    const createUrl = `http://${config.DOCKER_API_URL}/containers/create`;
    const createResponse = await axios.post(createUrl, data);
    const containerId = createResponse.data.Id;

    // Start the container
    const startUrl = `http://${config.DOCKER_API_URL}/containers/${containerId}/start`;
    await axios.post(startUrl);

    // The container will run and create the backup as specified in the command.

    // Optionally, you can wait for the container to finish if needed.
    // Wait for the container to finish (blocking)
    const waitUrl = `http://${config.DOCKER_API_URL}/containers/${containerId}/wait`;
    await axios.post(waitUrl);
}

async function createVolume(volumeName: string, type: string, path: PathType, username?: string, password?: string) {
    let data: any;

    switch (type) {
        case PathType.Local:
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
        case PathType.Nfs:
            data = {};
            break;
        case PathType.Cifs:
            data = {};
            break;
        default:
            data = {};
            break;
    }

    const response = await axios.post(`http://${config.DOCKER_API_URL}/volumes/create`, data);

    if (response.status !== 201) {
        throw new Error(response.data);
    }
}

export async function deleteVolume(volumeName) {
    const volumeToDelete = (await axios.get(`http://${config.DOCKER_API_URL}/volumes`)).data.Volumes.filter((volume: any) => volume.Name === volumeName);

    if (!volumeToDelete[0].Labels["com.dvb.volume"]) {
        throw new Error("Cannot delte volume. Volume was created outside of dvb");
    }

    const response = await axios.delete(`http://${config.DOCKER_API_URL}/volumes/${volumeName}`);

    if (response.status !== 204) {
        throw new Error(response.data);
    }
}

