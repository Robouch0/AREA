"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"

export async function login(emailValue: string, passwordValue: string) : Promise<boolean> {
    try {
        const response = await axiosInstance.post(`login/`, {
            email: emailValue,
            password: passwordValue
        });
        console.log(response);
        console.log(response.data);
        const cookiesObj = await cookies();
        cookiesObj.set('token', response.data);
        // cookies().set('token', response.data, {expires: 7, sameSite: 'Lax', secure: false});
        return true;
    } catch (error) {
        throw error;
    }
}

export async function checkAuthentification(token:string|undefined) {
    try {
        const response = await axiosInstance.get(`ping`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        console.log(response.data);
        return true;
    } catch (error) {
        console.log("Authentication check failed:", error);
        return false;
    }
}
