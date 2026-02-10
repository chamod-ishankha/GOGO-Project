import axios from "axios";

const API_GATEWAY_URL = "http://localhost:3000/api/v1/gogo"; // your backend URL

const api = axios.create({
    baseURL: API_GATEWAY_URL,
    headers: {
        "Content-Type": "application/json",
    },
});

// Add JWT interceptor
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("token");
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

export default api;
