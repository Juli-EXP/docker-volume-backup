// Import libraries
import axios from "axios";

//Import global variables
import { config } from "../config/variables";


// Return a list of all Docker volumes
export async function getVolumes() {
    const url = `http://${config.DOCKER_API_URL}/system/df`
    const volumesJson = (await axios.get(url)).data;

    const volumes = {
        Volumes: volumesJson.Volumes.map((volume: any) => ({
            Name: volume.Name,
            Size: volume.UsageData.Size,
            type: volume.Options ? volume.Options.type || "local" : "local",    // Checks if volume.Options and volume.Options.type exist. If not it sets the type to local
        })),
    };

    return volumes;
}

// Return a single Docker volume
export async function getVolume(dockerVolumeName: string) {
    const volumes = await getVolumes();

    const filteredVolumes = volumes.Volumes.filter((volume: any) => volume.Name === dockerVolumeName);

    return {
        Volumes: filteredVolumes,
    };

}

// Run a Docker container to backup a certain volume
async function backupVolume(volumeName: string, backupMount: string) {
    const data = {
        Image: `${config.CONTAINER_IMAGE}`,
        Cmd: ['ash', '-c', `tar cf /dest/${volumeName}.backup.tar /data/.`],
        HostConfig: {
            Binds: [`${volumeName}:/data`, `${backupMount}:/dest`],
            AutoRemove: true,
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