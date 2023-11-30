import axios from "axios"
import { error } from "console";
import { request } from "http";

const baseURL = process.env.NEXT_PUBLIC_API_URL;

const headers = {
    "Content-Type" : "application/json"
};

export const ApiClient = axios.create({baseURL, headers});

ApiClient.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        console.log(error)
        switch (error?.response?.status) {
            case 401:
                break
            case 404:
                break
            default:
                console.log("== internal server error");
        }
        const errorMessage = (error.response?.data?.message || "").split(",");
        throw new Error(errorMessage);
    }
)