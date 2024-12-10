import { FaGithub } from "react-icons/fa";
import { Button } from "@/components/ui/button";
import axios from "axios";
import { useEffect } from 'react';
import { useRouter} from "next/navigation";
import axiosInstance from "@/lib/axios";
async function redirectToGitHub() {
    try {
        const response = await axiosInstance.get(`oauth/github`);
        window.location.href = response.data;
    } catch (error) {
        console.error(error);
    }
}

async function askForToken(paramValue) {
    try {
        const response = await axiosInstance.post(`oauth/`, {
            service: "github",
            code: paramValue
        });

        await axios.post('/api/setToken', { token: response.data });
        return true;
    } catch (error) {
        console.error(error);
    }
}

export function GithubOauth() {
    const router = useRouter();
    useEffect(() => {
        const url = new URL(window.location.href);
        const paramValue = url.searchParams.get('code');

        if (paramValue) {
            askForToken(paramValue)
            .then(() => router.push("/services"))
            .catch((error) => console.log(error));
        }
    }, [router]);

    return (
        <Button
            className="focus-visible:border-slate-500 focus-visible:border-8 flex items-center justify-start px-6 bg-black hover:bg-black hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
            onClick={redirectToGitHub}
        >
            <FaGithub className="w-12 h-12"/>
            <p className="mx-3 text-2xl font-semibold">Continuer avec Github</p>
        </Button>
    );
}
