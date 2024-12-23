import { z, ZodType } from "zod";

// Types used for the REGISTER form verification
export type RegisterFormData = {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    confirmPassword: string;
};

export type RegisterValidFieldNames =
    | "firstName"
    | "lastName"
    | "email"
    | "password"
    | "confirmPassword"

export const RegisterUserSchema: ZodType<RegisterFormData> = z
    .object({
        firstName: z.string().min(1, { message: "First name is required" }),
        lastName: z.string().min(1, { message: "Last name is required" }),
        email: z.string().email({ message: "Invalid email address" }),
        password: z
            .string()
            .min(6, { message: "Password must be at least 6 characters long" }),
        confirmPassword: z.string()
    })
    .refine((data) : boolean => data.password === data.confirmPassword, {
        message: "Passwords don't match",
        path: ["confirmPassword"],
    });
