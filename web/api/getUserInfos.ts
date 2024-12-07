"use server";

import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"

export async function getUserInfo() {
    const cookiesObj = await cookies();
    const uid = cookiesObj.get("UID")?.value;

    try {
        const response = await axiosInstance.get(`users/${uid}`);
        console.log(response.data);
        return response.data;
    } catch (error) {
        throw error;
    }
}
