"use server";
import { redirect } from 'next/navigation';
import { checkAuthentification } from "@/api/authentification";

import { cookies } from 'next/headers';
import LoginForm from "@/components/pages/login/LoginForm";

export default async function LoginPage() {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (token != undefined) {
        const isAuthenticated = await checkAuthentification(token);

        if (isAuthenticated) {
            redirect('/services/myareas');
        }
    }
    return <LoginForm></LoginForm>;
}
