"use server";

import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";

export interface userInfo {
    id: number;
    first_name: string;
    last_name: string;
    email: string;
    password: string;
    providers: string[];
}
export async function getUserInfo(): Promise<userInfo> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookiesObj.get("UID")?.value;
        const token: string | undefined = cookiesObj.get('token')?.value;

        const response = await axiosInstance.get(`users/me`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        return response.data;
    } catch (error) {
        throw error;
    }
}

export async function getUserTokens(): Promise<string[]> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const arrTokens: string[] = [];
        const token: string | undefined = cookiesObj.get('token')?.value;

        const response = await axiosInstance.get(`/token/`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        if (response.data != null) {
            for (const token of response.data) {
                arrTokens.push(token.provider);
            }
        }
        return arrTokens;
    } catch (error) {
        console.error("Error fetching user tokens:", error);
        throw error;
    }
}
