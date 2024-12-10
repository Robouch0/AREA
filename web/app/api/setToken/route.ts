"use server";
import {NextRequest, NextResponse} from 'next/server';
import {cookies} from "next/headers";


export async function POST(request: NextRequest) {
    const { token } = await request.json();
    const response = NextResponse.json({ message: 'Token set successfully' });
    const cookiesObj = await cookies();
    cookiesObj.set("token", token);

    return response;
}
