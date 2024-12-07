"use client";

import Image from "next/image";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {FaEye, FaEyeSlash} from "react-icons/fa";
import {useState} from "react";

export default function ProfilePage(userData: any) {
    console.log(userData.userData.email);
    const [showPassword, setShowPassword] = useState(false);

    return (
        <>
            <div className="w-full h-full flex flex-col bg-white justify-center items-center my-16">
                <h1 className="text-6xl font-bold p-8"> Account Settings </h1>
                <hr className="w-1/3  my-8 bg-slate-800 opacity-20 h-1 border-0 dark:bg-gray-700"/>
                <div className="w-4/5 h-full bg-slate-900 rounded-3xl flex flex-col sm:flex-row p-8 m-2 justify-between gap-8">
                    <div
                        className="p-6 rounded-3xl w-full sm:w-1/2 bg-slate-700 h-full flex flex-col justify-center items-center gap-4">
                        <Image
                            className="border-4 rounded-b-full rounded-t-full border-black my-6"
                            alt="profile picture"
                            src="/default-avatar.png"
                            width="80"
                            height="80"
                        />
                        <h2 className="text-white text-2xl font-bold my-2 "> Email </h2>
                        <Input
                            type="email"
                            id="mail"
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="Email"
                            value={userData.userData.email}
                            disabled
                        />
                        <h2 className=" text-white text-2xl font-bold my-2"> First name </h2>
                        <Input
                            type="email"
                            id="mail"
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="text"
                            value={userData.userData.first_name}
                            disabled
                        />
                        <h2 className="p-2 text-white text-2xl font-bold"> Last name </h2>
                        <Input
                            type="email"
                            id="mail"
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="text"
                            value={userData.userData.last_name}
                            disabled
                        />
                        <h2 className="p-2 text-white text-2xl font-bold"> Password </h2>
                        <div className="w-full flex flex-row justify-center relative mb-6">
                            <Input
                                type={showPassword ? "text" : "password"}
                                id="password"
                                className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                aria-label="text"
                                value={userData.userData.password}
                                disabled
                            />
                            <Button
                                type="button"
                                onClick={() => setShowPassword(!showPassword)}
                                className="absolute top-1/2 right-32 transform -translate-y-1/2 bg-transparent border-none outline-none focus:outline-none hover:bg-transparent ring-0 shadow-none p-2"
                                aria-label={showPassword ? "Hide password" : "Show password"}
                            >
                                {showPassword ? <FaEyeSlash className="text-gray-500 scale-x-[-1] text-2xl"/> :
                                    <FaEye className="text-gray-500 scale-x-[-1] text-2xl"/>}
                            </Button>
                        </div>
                    </div>

                    <div
                        className="p-6 rounded-3xl w-full sm:w-1/2 bg-slate-700 h-full flex flex-col justify-center items-center gap-4">
                        <p className="text-white">
                            Right part
                        </p>
                    </div>
                </div>

            </div>
        </>
    );
}
