import { useEffect } from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import { oauhLogin } from "@/api/authentification";
import { Button } from '@/components/ui/utils/Button';

interface IOAuthButton {
    className: string,
    service: string,
    ServiceIcon?: React.ReactNode,
}

async function redirectToService(service: string) {
    try {
        const response = await axiosInstance.get(`oauth/${service}`, {
            params: {
                "redirect_uri": "http://127.0.0.1:8081"
            }
        });
        window.location.href = response.data;
    } catch (error) {
        console.error(error);
    }
}

async function askForToken(service: string, code: string | null) {
    try {
        await oauhLogin({ service: service, code: code, redirect_uri: "http://127.0.0.1:8081" })

        return true;
    } catch (error) {
        console.error(error);
    }
}

export function OauthButton({ service, className, ServiceIcon }: IOAuthButton) {
    const serviceDisplayName = service.charAt(0).toUpperCase() + service.slice(1);
    const router = useRouter();

    useEffect(() => {
        const url = new URL(window.location.href);
        const code: string | null = url.searchParams.get('code');

        if (code) {
            console.log(service, code)
            askForToken(service, code)
                .then(() => router.push("/services"))
                .catch((error) => console.log(error));
        }
    }, [router]);

    return (
        <Button
            className={className}
            onClick={() => { redirectToService(service) }}
        >
            {ServiceIcon}
            <p className="mx-3 text-2xl font-semibold">Continuer avec {serviceDisplayName}</p>
        </Button>
    );
}
