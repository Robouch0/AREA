import {useCallback, useEffect, useState} from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import {connectOauth, oauthLogin} from "@/api/authentification";
import { Button } from '@/components/ui/utils/Button';
import redirectURI from '@/lib/redirectUri';

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
        return response.data;
    } catch (error) {
        console.error(error);
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


    const openPopup = useCallback(async (service : string) => {
        const url:string = await redirectToService(service)

            if (typeof window !== 'undefined') {
                window.open(url, "popup", "width=1200,height=800,left=400,top=700,popup=true");
            }
    }, []);

    useEffect(() => {
        if (code) {
            askForToken(service, code, login)
                .then(() => router.push("/services"))
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
