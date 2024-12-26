import {useCallback, useEffect, useState} from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import { oauhLogin } from "@/api/authentification";
import { Button } from '@/components/ui/utils/Button';
import redirectURI from '@/lib/redirectUri';

interface IOAuthButton {
    className: string,
    service: string,
    ServiceIcon?: React.ReactNode,
}

async function redirectToService(service: string) {
    try {
        const response = await axiosInstance.get(`oauth/${service}`, {
            params: {
                "redirect_uri": redirectURI
            }
        });
        return response.data;
    } catch (error) {
        console.error(error);
    }
}
// http://127.0.0.1:8081
async function askForToken(service: string, code: string | null) {
    try {
        if (redirectURI == undefined)
            throw Error("env variable redirectURI is undefined")
        await oauhLogin({ service: service, code: code, redirect_uri: redirectURI })

        return true;
    } catch (error) {
        console.error(error);
    }
}

export function OauthButton({ service, className, ServiceIcon }: IOAuthButton) {
    const serviceDisplayName = service.charAt(0).toUpperCase() + service.slice(1);
    const router = useRouter();
    const [popupWindow, setPopupWindow] = useState<Window|null>(null);
    const [code, setPopupCode] = useState<string|null>("");


    const openPopup = useCallback(async (service) => {
        const url:string = await redirectToService(service)

            if (typeof window !== 'undefined') {
                const newWindow = window.open(url, "popup", "width=900,height=700,left=400,top=700,popup=true");
                setPopupWindow(newWindow);
            }
    }, []);


    useEffect(() => {
        if (!popupWindow) {
            return;
        }

        const checkPopup = setInterval(() => {
            if (popupWindow.closed) {
                clearInterval(checkPopup);
                setPopupWindow(null);
                return;
            }

            try {
                console.log(popupWindow.location.href);
                if (popupWindow.location.href.includes("code")) {
                    const url = new URL(popupWindow.location.href);
                    const code: string|null = url.searchParams.get('code');
                    window.opener.postMessage({ code: `${code}` }, 'http://127.0.0.1:8081');

                    popupWindow.close();
                    clearInterval(checkPopup);
                    setPopupWindow(null);
                }
            } catch (e) {
                console.error("error:", e);
            }
        }, 1000);

    }, [popupWindow]);

    useEffect(() => {
        const url = new URL(window.location.href);
        const code: string | null = url.searchParams.get('code');

        if (code) {
            console.log(service, code)
            askForToken(service, code)
                .then(() => router.push("/services"))
                .catch((error) => console.log(error));
        }
    }, [router, service]);

    return (
        <Button
            className={className}
            onClick={() => { openPopup(service) }}
        >
            {ServiceIcon}
            <p className="mx-3 text-2xl font-semibold">Continuer avec {serviceDisplayName}</p>
        </Button>
    );
}
