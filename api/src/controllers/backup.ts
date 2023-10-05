// Import libraries
import axios from "axios";

//Import global variables
import { DOCKER_API_URL } from "../config/variables";


export async function getDockerVolumeNames() {
    let volumes: any

    try {
        volumes = await axios.get(`http://${DOCKER_API_URL}/volumes`);
    } catch (error) {
        throw new Error(error);
    }

    const volumeNames = volumes.data.Volumes.map(obj => obj.Name);
    return volumeNames;
}


