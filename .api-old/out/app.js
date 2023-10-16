"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// Import libraries
const express_1 = __importDefault(require("express"));
const body_parser_1 = __importDefault(require("body-parser"));
// Import router
const routes_1 = __importDefault(require("./routes"));
// Import util functions
const error_1 = require("./utils/error");
//Import config variables
const variables_1 = require("./config/variables");
const init_1 = require("./config/init");
// Initialize program 
(0, init_1.init)()
    .then(() => {
    // Initialize express
    const app = (0, express_1.default)();
    app.use(body_parser_1.default.json());
    app.use("/api", routes_1.default);
    app.listen(variables_1.config.PORT, () => {
        return console.log(`Express is listening at http://localhost:${variables_1.config.PORT}`);
    });
})
    .catch((error) => {
    console.error('Initialization error:', (0, error_1.printError)(error));
    process.exit(1); // Exit the application on error if needed
});
//# sourceMappingURL=app.js.map