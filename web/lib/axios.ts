import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: process.env["NEXT_PUBLIC_GATEWAY_URL"]
});

export default axiosInstance;
