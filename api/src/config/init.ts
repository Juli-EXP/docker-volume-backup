// Import libraries
import axios from "axios";

//Import global variables
import { config } from "../config/variables";

export async function init() {
    console.log("Initializing program");

    console.log(`Pulling image: ${config.CONTAINER_IMAGE}`);
    const response = await pullImage();
    if (response) {
        console.log(`Successfully pulld image: ${config.CONTAINER_IMAGE}`);
    } else {
        throw new Error(`There was an error while pulling image: ${config.CONTAINER_IMAGE}`);
    }
}

// Pull the Docker image used for creating the Backup
async function pullImage() {
    const pullUrl = `http://${config.DOCKER_API_URL}/images/create`;
    const params = {
        fromImage: `${config.CONTAINER_IMAGE}`,
        //tag: 'latest',
    };

    // Send the API request to pull the image
    const response = await axios.post(pullUrl, null, { params });
    //console.log(response)

    // Check the response for success or handle any errors
    return response.status === 200;
}