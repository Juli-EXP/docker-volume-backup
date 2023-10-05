// Import libraries
import express from "express";
import bodyParser from "body-parser"

// Import router
import router from "./routes";

//Import global variables
import { PORT } from "./config/variables";


const app = express();

app.use(bodyParser.json());

app.use("/api", router);

app.listen(PORT, () => {
  return console.log(`Express is listening at http://localhost:${PORT}`);
});