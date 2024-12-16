"use server";

import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"

export interface userInfo {
    id: number;
    first_name: string;
    last_name: string;
    email: string;
    password: string;
}

export async function getUserInfo(): Promise<userInfo> {
    try {
        const cookiesObj = await cookies();
        const uid = cookiesObj.get("UID")?.value;
        console.log(uid);
        const response = await axiosInstance.get(`users/${uid}`);
        console.log(response.data);
        return response.data;
    } catch (error) {
        throw error;
    }
}
