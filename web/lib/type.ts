import { FieldError, UseFormRegister } from "react-hook-form";
import { z, ZodType } from "zod";

// types used for the LOGIN form verification
export type FormData = {
    email: string;
    password: string;
};

export type FormFieldProps = {
    type: string;
    placeholder: string;
    name: ValidFieldNames;
    register: UseFormRegister<FormData>;
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
            .min(4, { message: "Password is too short" })
    });
