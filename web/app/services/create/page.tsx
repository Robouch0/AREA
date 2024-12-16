"use server";

import { listAreas } from "@/api/listAreas";
import CreatePage from "@/components/ui/createPage";
import axiosInstance from "@/lib/axios";
import {cookies} from "next/headers";

export default async function Create() {
    try {
        const cookieStore = await cookies();
        const token = cookieStore.get('token')?.value;
        const uid = cookieStore.get("UID")?.value;

        const services = await listAreas()

        return <CreatePage {...services} uid></CreatePage>
    } catch (error) {
        console.log(error);
    }
}
