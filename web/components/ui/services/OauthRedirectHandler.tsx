'use client';

import {useEffect} from 'react';

export default function OAuthRedirectHandler() {

    useEffect(() => {
        const url = new URL(window.location.href);
        const searchParams = new URLSearchParams(url.search);
        const code: string | null = searchParams.get('code');
        const service: string | null = searchParams.get('state');

        if (window.opener && code && service) {
            try {
                window.opener.postMessage({type: "message", code: `${code},${service}`}, "/");
                setTimeout(() => {
                    console.log("Redirecting to main site");
                    window.close();
                }, 1200);
            } catch (error) {
                console.error("Error parsing state:", error);
            }
        }
    }, []);

    return (
    <>
        <div
            className="text-black bg-green-700 w-full h-24 flex flex-col justify-center items-center animate-pulse ease-in-out"
        >
            <p className="font-mono md:text-4xl text-xl font-bold"> Authentification successful, redirection you
                toward
                AREA </p>
        </div>
    </>)
}
