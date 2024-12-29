'use client';

import {useEffect} from 'react';

export default function OAuthRedirectHandler() {

    useEffect(() => {
        const url = new URL(window.location.href);
        const searchParams = new URLSearchParams(url.search);
        const service:string|null = searchParams.get("service");
        const code:string|null = searchParams.get('code');
        if (window.opener && code) {
            console.log(service);
            window.opener.postMessage({type: "message", code: `${code},${service}`}, "/");
            setTimeout(() => {
                console.log("Redirecting to main site");
                window.close();
            }, 1200);
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
