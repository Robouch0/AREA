'use server';


import ServicesLayout from "@/components/ui/servicesLayout";
import {cookies} from "next/headers";
import {checkAuthentification} from "@/api/authentification";
import {redirect} from 'next/navigation'
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";

export default async function InAppLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {

    const cookieStore:ReadonlyRequestCookies = await cookies();
    const token:string|undefined = cookieStore.get('token')?.value;
    const isAuthenticated:boolean = await checkAuthentification(token);

    if (!isAuthenticated) {
        redirect('/');
    }

    return (
        <>
            <ServicesLayout>
                <main className="grow">
                    {children}
                </main>
            </ServicesLayout>
        </>
    )
}
