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
        const data =  response.data.split(',');
        cookiesObj.set('token', data.at(0));
        cookiesObj.set('UID', data.at(1));
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
        console.info("Authentication check failed:", error);
        return false;
    }
}

export async function signUp(emailValue: string, passwordValue: string, firstNameValue: string, lastNameValue: string) : Promise<boolean> {
    try {
        const response = await axiosInstance.post(`sign-up/`, {
            email: emailValue,
            password: passwordValue,
            first_name: firstNameValue,
            last_name: lastNameValue
        });
        const loginResponse = await login(emailValue, passwordValue);
        console.log(loginResponse);
        return true;
    } catch (error) {
        throw error;
    }
}
