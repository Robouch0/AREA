"use server";

import { listAreas } from "@/api/listAreas";
import CreatePage from "@/components/pages/create/CreatePage";
import {cookies} from "next/headers";
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";
import {AreaServices} from "@/api/types/areaStatus";
import {getOauthProviders, getUserTokens} from "@/api/getUserInfos";

export interface TokenState {
    providerName: string;
    isTokenPresent: boolean;
}

function createTokenProviderState(existingProvider:string [], userTokens: string[] ) : TokenState[] {
    const oauthTokenStates: TokenState[] = []
    existingProvider.forEach((provider) => {
        let found: boolean = false;
        userTokens.forEach((token) => {
            if (token == provider) {
                found = true;
                oauthTokenStates.push({providerName: provider, isTokenPresent: found})
            }
        })
        if (!found) {
            oauthTokenStates.push({providerName: provider, isTokenPresent: false})
        }
    })
    oauthTokenStates.push({providerName: "crypto", isTokenPresent: true})
    oauthTokenStates.push({providerName: "dt", isTokenPresent: true})
    oauthTokenStates.push({providerName: "weather", isTokenPresent: true})
    return oauthTokenStates
}

export default async function Create() : Promise<React.JSX.Element|undefined> {
    try {
        const cookieStore: ReadonlyRequestCookies = await cookies();
        const uid: string | undefined = cookieStore.get("UID")?.value;

        const services : AreaServices[] = await listAreas()
        const existingProvider : string[] = await getOauthProviders()
        const userTokens : string[] = await getUserTokens()

        const oauthTokenStates : TokenState[] = createTokenProviderState(existingProvider, userTokens)

        if (uid == undefined) {
            throw Error("No User ID for the current user")
        }

        return <CreatePage services={services} userTokens={oauthTokenStates} uid={parseInt(uid)}/>
    } catch (error) {
        console.log(error);
    }
}
