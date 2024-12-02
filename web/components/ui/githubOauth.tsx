import {FaGithub} from "react-icons/fa";
import {Button} from "@/components/ui/button";

export function GithubOauth() {
    function redirectToGitHub() {
        const client_id = "Ov23linAJPU28i6OMP8G";
        const redirect_uri = "http://localhost:3000/services/";
        const scope = "read:user";

        const authUrl = `https://github.com/login/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&scope=${scope}`;


        window.location.href = authUrl;
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
