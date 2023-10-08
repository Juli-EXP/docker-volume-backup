// Import libraries
import express from "express";
import bodyParser from "body-parser"

// Import router
import router from "./routes";

// Import util functions
import { printError } from "./utils/error";

//Import config variables
import { config } from "./config/variables";
import { init } from "./config/init";


// Initialize program 
init()
  .then(() => {
    // Initialize express
    const app = express();

    app.use(bodyParser.json());

    app.use("/api", router);

    app.listen(config.PORT, () => {
      return console.log(`Express is listening at http://localhost:${config.PORT}`);
    });
  })
  .catch((error) => {
    console.error('Initialization error:', printError(error));
    process.exit(1); // Exit the application on error if needed
  });

