"use server"

import axiosInstance from "@/lib/axios";
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";
import {cookies} from "next/headers";

export async function deleteArea(areaId: number) : Promise<void> {
    try {
        const cookieStore : ReadonlyRequestCookies = await cookies();
        const token : string|undefined = cookieStore.get('token')?.value;

        await axiosInstance.delete(`areas/`, {
            headers: {
                'Authorization': `Bearer ${token}`
            },
            data: {
                area_id: areaId
            }
        });
    } catch (error) {
        throw error;
    }
}
