"use server";

import { listAreas } from "@/api/listAreas";
import CreatePage from "@/components/ui/createPage";
import {cookies} from "next/headers";

export default async function Create() {
    try {
        const cookieStore = await cookies();
        const uid = cookieStore.get("UID")?.value;

        const services = await listAreas()

        if (uid == undefined) {
            throw Error("No User ID for the current user")
        }

        return <CreatePage services={services} uid={parseInt(uid)}/>
    } catch (error) {
        console.log(error);
    }
}
