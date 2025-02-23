"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"

interface ActivateAreaRequest {
    area_id: number;
    activated: boolean;
}

export const activateArea = async (areaId: number, activated: boolean): Promise<void> => {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;

        if (token == undefined) {
            throw Error("Token is undefined");
        }
        console.log(areaId)
        const response = await axiosInstance.put<void>("/areas/activate",
            {
                area_id: areaId,
                activated: activated
            } as ActivateAreaRequest,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        );

        return response.data;
    } catch (err) {
        console.error("Error activating area:", err);
        throw err;
    }
};
