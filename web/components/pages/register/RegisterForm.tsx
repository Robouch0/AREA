"use client";
import { useForm } from "react-hook-form";
import { RegisterFormData, RegisterUserSchema } from "@/lib/typeRegister";
import { zodResolver } from "@hookform/resolvers/zod";
import { FaEye, FaEyeSlash, FaFacebook, FaGithub, FaGoogle } from "react-icons/fa";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { signUp } from "@/api/authentification";
import { OauthButton } from "@/components/ui/services/OauthButton";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { Button } from "@/components/ui/utils/Button";
import FormField from "@/components/ui/utils/FormField";

function RegisterForm() {
    const router: AppRouterInstance = useRouter();
    const [accountCreated, setAccountCreated] = useState(false);

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<RegisterFormData>({
        resolver: zodResolver(RegisterUserSchema),
    });
    useEffect(() : void => {
        if (accountCreated) {
            setTimeout((): void => {
                setAccountCreated(false);
                router.push('/services')
            }, 2300);
        }
    }, [accountCreated, router]);
    const [showPassword, setShowPassword] = useState(false);

    const onSubmit = async (data: RegisterFormData) => {
        try {
            const connected : boolean = await signUp(data.email, data.password, data.firstName, data.lastName);
            console.log(connected);
            setAccountCreated(true);
        } catch (error) {
            console.info(error);
        }
    };

    return (
        <>
            {accountCreated ?
                <div className="text-black bg-green-700 w-full h-24 flex flex-col justify-center items-center animate-pulse ease-in-out">
                    <p className="font-mono md:text-4xl text-xl font-bold"> Votre compte a été créé, nous vous redirigeons vers AREA </p>
                </div>
                : <> </>}
            <div className="flex items-center justify-center min-h-screen bg-white">
                <div className="flex flex-col items-center w-full max-w-md">
                    <h1 className="font-mono text-5xl font-extrabold mb-16 mt-8"> AREA </h1>
                    <h2 className="font-mono text-4xl font-black mb-12"> Sign up </h2>
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="grid col-auto">
                            <div className="w-full mb-6">
                                <FormField<RegisterFormData>
                                    type="text"
                                    placeholder="First name"
                                    name="firstName"
                                    register={register}
                                    error={errors.firstName}
                                    ariaLabel="first name inputfield"
                                    className="!text-2xl rounded-2xl font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                />
                            </div>
                            <div className="w-full mb-6">

                                <FormField<RegisterFormData>
                                    type="text"
                                    placeholder="Last name"
                                    name="lastName"
                                    register={register}
                                    error={errors.lastName}
                                    ariaLabel="last name inputfield"
                                    className="!text-2xl rounded-2xl font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                />
                            </div>
                            <div className="w-full mb-6">
                                <FormField<RegisterFormData>
                                    type="email"
                                    placeholder="Email"
                                    name="email"
                                    register={register}
                                    error={errors.email}
                                    ariaLabel="email inputfield"
                                    className="!text-2xl rounded-2xl font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                />
                            </div>
                            <div className="w-full mb-6 relative">

                                <FormField<RegisterFormData>
                                    type={showPassword ? "text" : "password"}
                                    placeholder="Password"
                                    name="password"
                                    ariaLabel="enter your password"
                                    register={register}
                                    error={errors.password}
                                    className="!text-2xl rounded-2xl font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                />
                                <Button
                                    type="button"
                                    onClick={(): void => setShowPassword(!showPassword)}
                                    className="absolute top-1/2 right-4 transform -translate-y-1/2 bg-transparent border-none outline-none focus-visible:outline-none hover:bg-transparent ring-0 shadow-none p-2"
                                    aria-label={showPassword ? "Hide password" : "Show password"}
                                >
                                    {showPassword ? <FaEyeSlash className="text-gray-500 scale-x-[-1] text-2xl" /> :
                                        <FaEye className="text-gray-500 scale-x-[-1] text-2xl" />}
                                </Button>
                            </div>
                            <div className="w-full mb-6 relative">

                                <FormField<RegisterFormData>
                                    type={showPassword ? "text" : "password"}
                                    placeholder="Confirm password"
                                    name="confirmPassword"
                                    ariaLabel="confirm your password"
                                    register={register}
                                    error={errors.confirmPassword}
                                    className="!text-2xl rounded-2xl font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                />
                                <Button
                                    type="button"
                                    onClick={() : void => setShowPassword(!showPassword)}
                                    className="absolute top-1/2 right-4 transform -translate-y-1/2 bg-transparent border-none outline-none focus-visible:outline-none hover:bg-transparent ring-0 shadow-none p-2"
                                    aria-label={showPassword ? "Hide password" : "Show password"}
                                >
                                    {showPassword ? <FaEyeSlash className="text-gray-500 scale-x-[-1] text-2xl" /> :
                                        <FaEye className="text-gray-500 scale-x-[-1] text-2xl" />}
                                </Button>
                            </div>

                            <div className="">
                                <Button
                                    type="submit"
                                    className="focus-visible:border-slate-500 focus-visible:border-8 rounded-full mt-8 w-full h-16 text-2xl font-bold"
                                    aria-label="Log In"
                                >
                                    Sign up
                                </Button>
                            </div>
                            <div className="inline-flex items-center justify-center w-full">
                                <hr className="w-2/3 h-px my-8 bg-black border-0 dark:bg-gray-700" />
                                <div
                                    className="absolute px-3 text-gray-400 -translate-x-1/2 bg-white left-1/2 dark:text-white dark:bg-gray-900">or
                                </div>
                            </div>
                            <div className="max-w-md w-full space-y-4">
                                <Button
                                    className="focus-visible:border-slate-500  focus-visible:border-8 flex items-center justify-start px-6 bg-blue-800 hover:bg-blue-800 hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
                                    type="button"
                                    arial-label="Facebook"
                                >
                                    <FaFacebook className="w-12 h-12" />
                                    <p className=" mx-3 text-2xl font-semibold"> Continuer avec Facebook </p>
                                </Button>

                                <OauthButton
                                    arial-label="Google"
                                    service="google"
                                    className="focus-visible:border-slate-500 focus-visible:border-8 flex items-center justify-start px-6 bg-red-500 hover:bg-red-500 hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
                                    ServiceIcon={<FaGoogle className="w-12 h-12"/>}
                                />      

                                <OauthButton
                                    arial-label="Github"
                                    service="github"
                                    className="focus-visible:border-slate-500 focus-visible:border-8 flex items-center justify-start px-6 bg-black hover:bg-black hover:opacity-90 rounded-3xl shadow-none h-20 w-full"
                                    ServiceIcon={<FaGithub className="w-12 h-12"/>}
                                />

                                <div className="flex flex-row font-bold">
                                    <p className="mb-8">
                                        Vous avez déjà un compte ?
                                    </p>
                                    <button
                                        className="mb-8 mx-2 underline-offset-1 underline font-bold hover:cursor-pointer focus-visible:border-4 focus-visible:border-slate-700 focus-visible:outline-none focus-visible:p-2 rounded-3xl"
                                        onClick={() : void => router.push('/')}
                                        tabIndex={0}
                                    >
                                        Connectez-vous ICI !
                                    </button>
                                </div>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </>
    );
}

export default RegisterForm;
