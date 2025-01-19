import { useCallback, useEffect, useState } from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import { connectOauth, oauthLogin } from "@/api/authentification";
import { Button } from '@/components/ui/utils/Button';
import redirectURI from '@/lib/redirectUri';
import { useToast } from "@/hooks/use-toast"
import * as React from "react";

interface IOAuthButton {
    className: string,
    service: string,
    ServiceIcon?: React.ReactNode | null,
    textButton: string,
    location: string,
}

async function redirectToService(service: string) {
    try {
        const response = await axiosInstance.get(`oauth/${service}`, {
            params: {
                "redirect_uri": redirectURI,
            }
        });
        return response.data;
    } catch (error) {
        throw error;
    }
}
// http://127.0.0.1:8081
async function askForToken(service: string, code: string | null, location: string) {

    try {
        if (redirectURI == undefined)
            throw Error("env variable redirectURI is undefined")
        console.log(location)
        if (location == "login") {
            await oauthLogin({ service: service, code: code, redirect_uri: redirectURI })
        } else {
            await connectOauth(service, code)
        }
        return true;
    } catch (error) {
        console.error(error);
    }
}

export function OauthButton({ service, className, ServiceIcon, textButton, location}: IOAuthButton) {
    const router = useRouter();
    const [code, setPopupCode] = useState<string | null>("");
    const { toast } = useToast()
    const [channel, setChannel] = useState<BroadcastChannel | null>(null);
    console.log(service)
    useEffect(() => {
        const handleMessage = (event: MessageEvent) => {
            if (!event || !event.data.code) {
                return;
            }
            const eventValues: string[] = event.data.code.split(",")
            const code: string = eventValues[0];
            const msgService: string = eventValues[1];
            if (event.data.type === "message" && msgService === service) {
                setPopupCode(code);
            }
        }

        if (channel == null) {
            return
        }
        channel.onmessage = handleMessage
    }, [channel, service])

    const openPopup = useCallback(async (service: string) => {
        try {
            const url = await redirectToService(service);
            if (typeof window !== 'undefined' && url !== "") {
                window.open(url, "popup", "width=1200,height=800,left=400,top=700,popup=true");
            }
            setChannel(new BroadcastChannel(`oauth-${service}`))
        } catch (err) {
            toast({
                title: "ERROR : Area server are down for the moment",
                description: "Failed to get the Oauth provider URL. Please try again later.",
                variant: 'destructive',
                duration: 2000,
            });
            console.log("ERROR trying to reach server" + err);
        }
    }, [toast]);

    useEffect(() => {
        if (code) {
            askForToken(service, code, location)
                .then(() => {
                    if (location == "login") {
                        router.push("/services/myareas")
                    } else if (location == "profile") {
                        console.log("ici")
                        router.push("/services/profile")
                    } else if (location == "create") {
                        window.location.reload()
                    }
                })
                .catch((error) => console.log(error));
        }
    }, [code, router, service, location]);

    return (
        <Button
            className={className}
            onClick={() => { openPopup(service) }}
        >
            {ServiceIcon == null ?
                <></>
                :
                <div className={"mx-3"}>
                    {ServiceIcon}
                </div>
            }
            <p className={"text-black font-bold text-xl"}>{textButton}</p>
        </Button>
    );
}
