"use server";
import {cookies} from 'next/headers';
import axiosInstance from "@/lib/axios"
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";

export async function unlinkOauthProvider(provider: string): Promise<void> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookiesObj.get("UID")?.value;
        const token : string|undefined = cookiesObj.get('token')?.value;

        if (!uid) {
            throw new Error("User ID not found in cookies");
        }
        await axiosInstance.delete(`token/${provider}`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

    } catch (error) {
        console.error("Error in unlinkOauth:", error);
        throw error;
    }
}
