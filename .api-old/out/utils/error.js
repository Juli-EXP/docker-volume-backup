"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.printError = void 0;
function printError(error) {
    return {
        status: error.response.status,
        statusText: error.response.statusText,
        message: error.response.data.message || "",
    };
}
exports.printError = printError;
//# sourceMappingURL=error.js.map