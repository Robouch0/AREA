"use server";
import { UserSchema } from "@/lib/type";
import { NextResponse } from "next/server";

export async function POST(request: Request) {
    const body = await request.json();
    const result = UserSchema.safeParse(body);

    if (result.success) {
        return NextResponse.json({ success: true });
    }

    const serverErrors = Object.fromEntries(
        result.error?.issues?.map((issue) => [issue.path[0], issue.message]) || []
    );

    return NextResponse.json({ errors: serverErrors });
}
