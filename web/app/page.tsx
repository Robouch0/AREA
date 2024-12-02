'use client';
import {Button} from "@/components/ui/button";
import {Input} from "@/components/ui/input";
import {useState} from "react";
import {FaEye, FaEyeSlash} from "react-icons/fa";
import { redirect } from 'next/navigation'
import {FaGoogle, FaFacebook, FaGithub} from 'react-icons/fa';
import {GithubOauth} from "@/components/ui/githubOauth";

export default function Page() {
    const [showPassword, setShowPassword] = useState(false);
    const [emailValue, setEmailValue] = useState('');
    const [passwordValue, setPasswordValue] = useState('');

    const handleSubmit = (event:React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        redirect('/services');
    };

    const handleGoogleOauth = () => {

    };

    const togglePasswordVisibility = () => {
        setShowPassword(!showPassword);
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-white">
            <div className="flex flex-col items-center w-full max-w-md">
                <h1 className="font-mono text-5xl font-extrabold mb-16"> AREA </h1>
                <h2 className="font-mono text-4xl font-black mb-12"> Log In </h2>
                <form id="loginForm" onSubmit={handleSubmit} className="">
                    <div className="w-full mb-6">
                        <Input
                            type="email"
                            placeholder="Email"
                            id="mail"
                            className="!text-2xl rounded-2xl font-extrabold border-4 focus:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="Email"
                            value={emailValue}
                            onChange={(e) => setEmailValue(e.target.value)}
                            required
                        />
                    </div>
                    <div className="w-full mb-6 relative">
                        <Input
                            type={showPassword ? "text" : "password"}
                            placeholder="Password"
                            id="password"
                            className="!text-2xl rounded-2xl border-4 focus:border-black w-full p-4 h-16 pr-12 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60 font-extrabold"
                            aria-label="Password"
                            value={passwordValue}
                            onChange={(e) => setPasswordValue(e.target.value)}
                            required
                        />
                        <Button
                            type="button"
                            onClick={togglePasswordVisibility}
                            className="absolute top-1/2 right-4 transform -translate-y-1/2 bg-transparent border-none outline-none focus:outline-none hover:bg-transparent ring-0 shadow-none p-2"
                            aria-label={showPassword ? "Hide password" : "Show password"}
                        >
                            {showPassword ? <FaEyeSlash className="text-gray-500 scale-x-[-1] text-2xl"/> :
                                <FaEye className="text-gray-500 scale-x-[-1] text-2xl"/>}
                        </Button>
                    </div>
                    <div className="">
                        <Button
                            type="submit"
                            className="rounded-full mt-8 w-full h-16 text-2xl font-bold"
                            aria-label="Log In"
                        >
                            Log In
                        </Button>
                    </div>
                </form>
                <div className="inline-flex items-center justify-center w-full">
                    <hr className="w-2/3 h-px my-8 bg-black border-0 dark:bg-gray-700"/>
                    <div
                        className="absolute px-3 text-gray-400 -translate-x-1/2 bg-white left-1/2 dark:text-white dark:bg-gray-900">or
                    </div>
                </div>
                <div className="max-w-md w-full space-y-4">
                    <Button
                        className="flex items-center justify-start px-6 bg-blue-800 hover:bg-blue-800 hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
                        type="button"
                        arial-label="Facebook"
                    >
                        <FaFacebook className="w-12 h-12"/>
                        <p className=" mx-3 text-2xl font-semibold"> Continuer avec Facebook </p>
                    </Button>

                    <Button
                        className="flex items-center justify-start px-6 bg-red-500 hover:bg-red-500 hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
                        type="button"
                        arial-label="Google"
                        onClick={handleGoogleOauth}
                    >
                        <FaGoogle className="w-12 h-12"/>
                        <p className=" mx-3 text-2xl font-semibold"> Continuer avec Google </p>
                    </Button>

                    <GithubOauth
                        arial-label="Github"
                    />
                    <div className="flex flex-row font-bold">
                        <p>
                            Vous n'avez pas encore de compte ?
                        </p>
                        <a
                            className="mx-2 underline-offset-1 underline font-bold hover:cursor-pointer"
                            onClick={() => redirect('/register')}
                        >
                            Inscrivez-vous ICI !
                        </a>
                    </div>
                </div>
            </div>
        </div>
    );
}
