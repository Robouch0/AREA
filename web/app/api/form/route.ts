"use server";
import {UserSchema} from "@/lib/typeLogin";
import {NextResponse} from "next/server";
import {ZodIssue} from "zod";

export async function POST(request: Request) {
    const body = await request.json();
    const result = UserSchema.safeParse(body);

    if (result.success) {
        return NextResponse.json({success: true});
    }

    const serverErrors = Object.fromEntries(
        result.error?.issues?.map((issue: ZodIssue): [string | number, string] => [issue.path[0], issue.message]) || []
    );

    return NextResponse.json({errors: serverErrors});
}
