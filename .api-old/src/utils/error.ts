/*
*   Error handling
*
*
*/

export enum ErrorType {
    Default = "default",
    Axios = "axios",
}

export function formatError(error: any, errorType?: ErrorType) {
    let formattedError: any;

    switch (errorType) {
        case ErrorType.Axios:
            formattedError = {
                status: error.response.status,
                statusText: error.response.statusText,
                message: error.response.data.message || "",
            };
            break;
    
        default:
            formattedError = error;
            break;
    }

    return formattedError;
}
