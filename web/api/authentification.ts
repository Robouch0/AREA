"use server";
import {cookies} from 'next/headers';
import axiosInstance from "@/lib/axios"
import {AxiosResponse} from 'axios';
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";
import RedirectUri from "@/lib/redirectUri";

export async function login(emailValue: string, passwordValue: string): Promise<boolean> {
    try {
        const response: AxiosResponse<UserLogInfosBody, string> = await axiosInstance.post<UserLogInfosBody, AxiosResponse<UserLogInfosBody, string>, UserCredentials>(
            `login/`, {
                email: emailValue,
                password: passwordValue
            });

        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const {token, user_id} = response.data
        cookiesObj.set('token', token);
        cookiesObj.set('UID', user_id.toString());
        return true;
    } catch (error) {
        throw error;
    }
}

export async function oauthLogin(oauthLogBody: OAuthLoginBody): Promise<void> {
    try {
        const response: AxiosResponse<UserLogInfosBody, string> = await axiosInstance.post<UserLogInfosBody, AxiosResponse<UserLogInfosBody, string>>(
            `oauth/`, oauthLogBody);

        const cookiesObj : ReadonlyRequestCookies = await cookies();
        const {token, user_id} = response.data

        cookiesObj.set('token', token);
        cookiesObj.set('UID', user_id.toString());
    } catch (error) {
        throw error
    }
}

export async function connectOauth(service: string, code: string|null): Promise<void> {
    try {
        const cookiesObj: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookiesObj.get("UID")?.value;
        const token : string|undefined = cookiesObj.get('token')?.value;

        if (!uid) {
            throw new Error("User ID not found in cookies");
        }
        const response = await axiosInstance.post("oauth/connect/   ", {
            service: service,
            code: code,
            redirect_uri: RedirectUri
        }, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        console.log(response);
        return response.data;
    } catch (error) {
        console.error("Error in connectOauth:", error);
        throw error;
    }
}

export async function checkAuthentification(token: string | undefined) : Promise<boolean> {
    try {
        const response = await axiosInstance.get(`ping`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        console.log(response.data);
        return true;
    } catch (error) {
        // console.info("Authentication check failed:", error);
        return false;
    }
}

export async function signUp(
    emailValue: string, passwordValue: string, firstNameValue: string, lastNameValue: string): Promise<boolean> {
    try {
        await axiosInstance.post(`sign-up/`, {
            email: emailValue,
            password: passwordValue,
            first_name: firstNameValue,
            last_name: lastNameValue
        });
        const loginResponse : boolean = await login(emailValue, passwordValue);
        console.log(loginResponse);
        return true;
    } catch (error) {
        throw error;
    }
}
