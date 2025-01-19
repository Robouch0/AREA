"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";

export async function updateUserInfos(firstName: string, lastName: string, password: string): Promise<void> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookiesObj.get("UID")?.value;
        const token: string | undefined = cookiesObj.get('token')?.value;
        console.log(password)
        if (!uid) {
            throw new Error("User ID not found in cookies");
        }
        await axiosInstance.put(`user/me`, {
            first_name: firstName,
            last_name: lastName,
            password: password
        }, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
    } catch (error) {
        console.error("Error in Update:", error);
        throw error;
    }
}
