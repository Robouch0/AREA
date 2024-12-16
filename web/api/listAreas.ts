"use server"

import axiosInstance from "@/lib/axios";
import {cookies} from "next/headers";

export const listAreas = async (): Promise<AreaServices[]> => {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;
        
        if (token == undefined) {
            throw Error("Token is undefined")
        }
        const response = await axiosInstance.get<AreaServices[]>(`create/list`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        return response.data
    } catch (err) {
        throw err
    }
}
