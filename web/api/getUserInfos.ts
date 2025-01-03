"use server";

import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";

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
        const uid: string|undefined = cookiesObj.get("UID")?.value;
        console.log(uid);
        const response = await axiosInstance.get(`users/${uid}`);
        console.log(response.data);
        return response.data;
    } catch (error) {
        throw error;
    }
}

export async function getUserTokens(): Promise<string[]> {
    try {
        const cookiesObj:ReadonlyRequestCookies = await cookies();
        const uid : string |undefined = cookiesObj.get("UID")?.value;
        const arrTokens : string[] = [];

        const response = await axiosInstance.get(`/token/${uid}`);
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
