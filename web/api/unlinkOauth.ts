"use server";
import {cookies} from 'next/headers';
import axiosInstance from "@/lib/axios"
import {AxiosResponse} from 'axios';
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";

export async function unlinkOauthProvider(provider: string): Promise<void> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookiesObj.get("UID")?.value;
        const token : string|undefined = cookiesObj.get('token')?.value;

        if (!uid) {
            throw new Error("User ID not found in cookies");
        }
        const response = await axiosInstance.post("oauth/connect/   ", {
            service: service,
            code: code,
        }, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        console.log(response);
        return response.data;
    } catch (error) {
        console.error("Error in connectOauth:", error);
        throw error;
    }
}
