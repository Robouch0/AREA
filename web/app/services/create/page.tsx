"use server";

import CreatePage from "@/components/ui/createPage";
import axiosInstance from "@/lib/axios";
import {cookies} from "next/headers";

export default async function Create() {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;
        const response = await axiosInstance.get(`create/list`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        const dataObject = response.data;
        return <CreatePage {...dataObject}></CreatePage>
    } catch (error) {
        console.log(error);
    }
}
