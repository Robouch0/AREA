"use server";

import { listAreas } from "@/api/listAreas";
import CreatePage from "@/components/pages/create/CreatePage";
import {cookies} from "next/headers";
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";
import {AreaServices} from "@/api/types/areaStatus";
import {getUserTokens} from "@/api/getUserInfos";

export default async function Create() : Promise<React.JSX.Element|undefined> {
    try {
        const cookieStore: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookieStore.get("UID")?.value;

        const services : AreaServices[] = await listAreas()
        const userTokens : string[] = await getUserTokens()

        if (uid == undefined) {
            throw Error("No User ID for the current user")
        }

        return <CreatePage services={services} userTokens={userTokens} uid={parseInt(uid)}/>
    } catch (error) {
        console.log(error);
    }
}
