import {useCallback, useEffect, useState} from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import {connectOauth, oauthLogin} from "@/api/authentification";
import { Button } from '@/components/ui/utils/Button';
import redirectURI from '@/lib/redirectUri';
import {useToast} from "@/hooks/use-toast"

interface IOAuthButton {
    className: string,
    service: string,
    ServiceIcon?: React.ReactNode|null,
    textButton:string,
    login: boolean,
}

async function redirectToService(service: string) {
    try {
        const response = await axiosInstance.get(`oauth/${service}`, {
            params: {
                "redirect_uri": redirectURI + "?service=" + service,
            }
        });
        console.log('response data: ' + response.data);
        return response.data;
    } catch (error) {
        throw error;
    }
}
// http://127.0.0.1:8081
async function askForToken(service: string, code: string | null, login:boolean) {

    try {
        if (redirectURI == undefined)
            throw Error("env variable redirectURI is undefined")
        console.log(login)
        if (login) {
            await oauthLogin({service: service, code: code, redirect_uri: redirectURI})
        } else {
            await connectOauth(service, code)
        }
        return true;
    } catch (error) {
        console.error(error);

    }
}

export function OauthButton({ service, className, ServiceIcon, textButton, login }: IOAuthButton) {
    const router = useRouter();
    const [code, setPopupCode] = useState<string|null>("");
    const { toast } = useToast()


    const openPopup = useCallback(async (service : string) => {
        try {
            const url = await redirectToService(service);
            if (typeof window !== 'undefined' && url !== "") {
                window.open(url, "popup", "width=1200,height=800,left=400,top=700,popup=true");
            }
        } catch (err) {
            toast({
                title: "ERROR : Area server are down for the moment",
                description: "Failed to get the Oauth provider URL. Please try again later.",
            });
            console.log("ERROR trying to reach server" + err);
        }
    }, [toast]);

    useEffect(() => {
        if (code) {
            askForToken(service, code, login)
                .then(() => {
                    if (login) {
                        router.push("/services")
                    } else {
                        router.push("/services/profile")
                    }
                })
                .catch((error) => console.log(error));
        }
    }, [code, router, service, login]);

    useEffect(()  => {
        const handleMessage = (event:MessageEvent) => {
            if (!event || !event.data.code) {
                return;
            }
            // maybe we will later need to check for message origin, for security purposes
            const eventValues:string[] = event.data.code.split(",")
            const code: string = eventValues[0];
            const msgService :string = eventValues[1];
            if (event.data.type === "message" && msgService === service) {
                setPopupCode(code);
            }
        };

        window.addEventListener('message', handleMessage);
        return () => window.removeEventListener('message', handleMessage);
    }, [service]);

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
            <p >{textButton}</p>
        </Button>
    );
}
