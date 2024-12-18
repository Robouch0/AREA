import { FaGithub } from "react-icons/fa";
import { Button } from "@/components/ui/utils/Button";
import { useEffect } from 'react';
import { useRouter } from "next/navigation";
import axiosInstance from "@/lib/axios";
import { oauhLogin } from "@/api/authentification";
import {AppRouterInstance} from "next/dist/shared/lib/app-router-context.shared-runtime";

async function redirectToGitHub() : Promise<void> {
    try {
        const response = await axiosInstance.get(`oauth/github`);
        window.location.href = response.data;
    } catch (error) {
        console.error(error);
    }
}

async function askForToken(code: string | null) : Promise<true|undefined> {
    try {
        await oauhLogin({ service: "github", code: code }) // Encore à voir si c bon !

        return true;
    } catch (error) {
        console.error(error);
    }
}

export function GithubOauth() {
    const router: AppRouterInstance = useRouter();
    useEffect(() : void => {
        const url = new URL(window.location.href);
        const paramValue: string | null = url.searchParams.get('code');

        if (paramValue) {
            askForToken(paramValue)
                .then(() => router.push("/services"))
                .catch((error) => console.log(error));
        }
    }, [router]);

    return (
        <Button
            className="focus-visible:border-slate-500 focus-visible:border-8 flex items-center justify-start px-6 bg-black hover:bg-black hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
            onClick={(): void => {redirectToGitHub()}}
        >
            <FaGithub className="w-12 h-12" />
            <p className="mx-3 text-2xl font-semibold">Continuer avec Github</p>
        </Button>
    );
}
