"use server";
import { cookies } from 'next/headers';
import axiosInstance from "@/lib/axios"
import {PossibleType} from "@/app/services/create/page";

axiosInstance.interceptors.request.use(request => {
    console.log('Starting Request', JSON.stringify(request, null, 2))
    return request
})
export async function create(data: Record<string, PossibleType>) {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;
        const response = await axiosInstance.post(`create/dt`, data, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });


        console.log(response);
    } catch (error) {
        console.log(error);
    }
}

export async function testCreateModifyRepo() {

    const createRequest: Record<string, PossibleType> = {};
    // if (params != undefined) {
    const action: Record<string, PossibleType> = {};
    const ingredientsA: Record<string, PossibleType> = {};
    const reaction: Record<string, PossibleType> = {};
    const ingredientsR: Record<string, PossibleType> = {};

    action["service"] = "dateTime";
    action["microservice"] = "dateTimeTrigger";
    ingredientsA["minutes"] = 12;
    ingredientsA["hours"] = 12;
    ingredientsA["day_month"] = 10;
    ingredientsA["month"] = 12;
    ingredientsA["day_week"] = 2;
    action["ingredients"] = ingredientsA;

    reaction["service"] = "github";
    reaction["microservice"] = "updateRepo";
    ingredientsR["owner"] = "Robouch0";
    ingredientsR["repo"] = "testWhanos";

    ingredientsR["name"] = "tkt";
    ingredientsR["description"] = "mdr";
    reaction["ingredients"] = ingredientsR;
    createRequest["user_id"] = 1;
    createRequest["action"] = action;
    createRequest["reaction"] = reaction;
    console.log(createRequest);
    try {
        const result = await create(createRequest);
        console.log('Create result:', result);
    } catch (error) {
        console.error('Error in create:', error);
    }
}
