"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";
import {AreaCreateBody} from "@/api/types/areaCreateBody";

// Axios debug interceptors
axiosInstance.interceptors.request.use(request => {
    // console.log('Starting Request', JSON.stringify(request, null, 2))
    return request
})
export async function create(data: AreaCreateBody) : Promise<void> {
    try {
        const cookieStore : ReadonlyRequestCookies = await cookies();
        const token : string|undefined = cookieStore.get('token')?.value;

        await axiosInstance.post(`areas/create/${data.action.service}`, data, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

    } catch (error) {
        throw error;
    }
}
