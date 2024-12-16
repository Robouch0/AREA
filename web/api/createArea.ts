"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"

axiosInstance.interceptors.request.use(request => {
    console.log('Starting Request', JSON.stringify(request, null, 2))
    return request
})
export async function create(data: AreaCreateBody) {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;
        await axiosInstance.post(`create/dt`, data, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

    } catch (error) {
        console.log("ERROR");
    }
}
