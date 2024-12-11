import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: process.env.GATEWAY_URL
    // baseURL: 'http://localhost:3000/'
});

export default axiosInstance;
