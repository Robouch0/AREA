"use server";
import {NextRequest, NextResponse} from 'next/server';
import {cookies} from "next/headers";
import {ReadonlyRequestCookies} from "next/dist/server/web/spec-extension/adapters/request-cookies";


export async function POST(request: NextRequest) {
    const { token } = await request.json();
    const response : NextResponse<{message: string}> = NextResponse.json({ message: 'Token set successfully' });
    const cookiesObj : ReadonlyRequestCookies = await cookies();
    cookiesObj.set("token", token);

    return response;
}
