import {FaGithub} from "react-icons/fa";
import {Button} from "@/components/ui/button";
import axios from "axios";
async function redirectToGitHub() {
    try {
        const response = await axios.get(`http://localhost:3000/oauth/github`);
        console.log(response.data);
        window.location.href = response.data;
        return true;
    } catch (error) {
        throw error;
    }

}

async function askForToken(paramValue:string) {
    try {
        const response = await axios.post(`http://localhost:3000/oauth/`, {
            service : "github",
            code : paramValue
        });
        console.log(response.data);
        return true;
    } catch (error) {
        console.log(error);
    }
}

export function GithubOauth() {
    if (typeof window !== 'undefined') {
        const url = window.location.href;
        const params = new URLSearchParams(new URL(url).search);
        const paramValue = params.get('code');
        if (paramValue != null) {
            askForToken(paramValue);
        }
    }
    return (
        <>
            <Button
                className="flex items-center justify-start px-6 bg-black hover:bg-black hover:opacity-80 rounded-3xl shadow-none h-20 w-full"
                onClick={redirectToGitHub}
            >

                <FaGithub className="w-12 h-12"/>
                <p className=" mx-3 text-2xl font-semibold"> Continuer avec Github </p>
            </Button>
        </>
    );
}
