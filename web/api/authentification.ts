"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import { AxiosResponse } from 'axios';

export async function login(emailValue: string, passwordValue: string): Promise<boolean> {
    try {
        const response = await axiosInstance.post<UserLogInfosBody, AxiosResponse<UserLogInfosBody, string>, UserCredentials>(`login/`, {
            email: emailValue,
            password: passwordValue
        });

        const cookiesObj = await cookies();
        const { token, user_id } = response.data
        cookiesObj.set('token', token);
        cookiesObj.set('UID', user_id.toString());
        return true;
    } catch (error) {
        throw error;
    }
}

export async function oauhLogin(oauthLogBody: OAuthLoginBody) {
    try {
        const response = await axiosInstance.post<UserLogInfosBody, AxiosResponse<UserLogInfosBody, string>>(`oauth/`, oauthLogBody);
        
        const cookiesObj = await cookies();
        const { token, user_id } = response.data

        cookiesObj.set('token', token);
        cookiesObj.set('UID', user_id.toString());
    } catch (error) {
        throw error;
    }
}

export async function checkAuthentification(token: string | undefined) {
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

export async function signUp(emailValue: string, passwordValue: string, firstNameValue: string, lastNameValue: string): Promise<boolean> {
    try {
        await axiosInstance.post(`sign-up/`, {
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
