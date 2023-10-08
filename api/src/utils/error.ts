export function printError(error: any) {
    return {
        status: error.response.status,
        statusText: error.response.statusText,
        message: error.response.data.message || "",
    };
}