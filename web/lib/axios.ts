import axios, {AxiosInstance} from 'axios';

const axiosInstance : AxiosInstance = axios.create({
    baseURL: process.env["NEXT_PUBLIC_GATEWAY_URL"]
});

export default axiosInstance;
