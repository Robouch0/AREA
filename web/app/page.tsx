"use server";
import { redirect } from 'next/navigation';
import { checkAuthentification } from "@/api/authentification";
import Login from '@/components/ui/login';

import { cookies } from 'next/headers';

export default async function LoginPage() {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    const isAuthenticated = await checkAuthentification(token);

    if (isAuthenticated) {
        redirect('/services');
    }

    return <Login />;
}
