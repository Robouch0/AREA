import { FieldError, UseFormRegister } from "react-hook-form";
import { z, ZodType } from "zod";

// types used for the LOGIN form verification
export type FormData = {
    email: string;
    password: string;
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type FormFieldProps<T extends Record<string, any>> = {
    type: string;
    placeholder: string;
    name: keyof T;
    register: UseFormRegister<T>;
    error: FieldError | undefined;
    valueAsNumber?: boolean;
    ariaLabel?: string;
    className?: string;
};


export type ValidFieldNames =
    | "email"
    | "password"

export const UserSchema: ZodType<FormData> = z
    .object({
        email: z.string().email(),
        password: z
            .string()
            .min(6, { message: "Password must be at least 6 characters long" })
    });
