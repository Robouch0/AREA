"use client";

import Image from "next/image";
import {Input} from "@/components/ui/utils/Input";
import {Button} from "@/components/ui/utils/Button";
import {FaEye, FaEyeSlash} from "react-icons/fa";
import {ScrollArea} from "@/components/ui/utils/Scroll-area"
import {ChangeEvent, useState} from "react";
import {ServiceIcon} from "@/components/ui/services/ServiceIcon";
import {userInfo} from "@/api/getUserInfos";
import {OauthButton} from "@/components/ui/services/OauthButton";
import {unlinkOauthProvider} from "@/api/unlinkOauth";
import {updateUserInfos} from "@/api/updateUserInfos";
import {useToast} from "@/hooks/use-toast";




export default function ProfilePage({email, first_name, last_name, password, providers}: userInfo) {
    const [showPassword, setShowPassword] = useState(false);
    const [firstName, setFirstName] = useState(first_name);
    const [lastName, setLastName] = useState(last_name);
    const [passw, setPassword] = useState(password);
    const [passwTooShort, setTooShort] = useState(false);
    const { toast } = useToast()

    const tags : string[] = [
        "github",
        "google",
        "discord",
        "spotify",
        "gitlab",
        "asana",
        "miro",
    ]

    function handleDisconnectProvider(provider: string) {
        unlinkOauthProvider(provider);
    }
    function handleDataUpdate() {
        updateUserInfos(firstName, lastName, passw).then(() => {
        toast({
            title: "Update sucessful",
            description: "Your new datas have been updated on our server.",
            variant: 'default',
            duration: 2500,
        })}
        ).catch(() => {
        toast({
            title: "Update failed",
            description: "Your new datas have not been updated on our server.",
            variant: 'destructive',
            duration: 2500,
        })}
        )
    }

    const reloadPage = () => {
        window.location.reload()
    }

    return (
        <>
            <div className="pt-10 w-full h-full flex flex-col bg-white justify-center items-center my-16">
                <h1 className="text-6xl font-bold p-8"> Account Settings </h1>
                <hr className="w-1/3  my-8 bg-slate-800 opacity-20 h-1 border-0 dark:bg-gray-700"/>
                <div
                    className="w-4/5 h-full bg-slate-900 rounded-3xl flex flex-col lg:flex-row p-8 m-2 justify-between"
                >
                    <div
                        className="p-6 rounded-3xl w-full lg:w-1/2 bg-slate-700 h-full flex flex-col justify-center items-center gap-4"
                    >
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
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-8 focus-visible:ring-0 focus-visible:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="Email"
                            value={email}
                            disabled
                        />
                        <h2 className=" text-white text-2xl font-bold my-2"> First name </h2>
                        <Input
                            type="text"
                            id="firstname"
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-8 focus-visible:ring-0 focus-visible:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="text"
                            value={firstName}
                            onChange={(e:ChangeEvent<HTMLInputElement>) => {
                                setFirstName(e.target.value);
                            }}
                        />
                        <h2 className="p-2 text-white text-2xl font-bold"> Last name </h2>
                        <Input
                            type="text"
                            id="lastname"
                            className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-8 focus-visible:ring-0 focus-visible:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                            aria-label="text"
                            value={lastName}
                            onChange={(e:ChangeEvent<HTMLInputElement>) => {
                                setLastName(e.target.value);
                            }}
                        />
                        <h2 className="p-2 text-white text-2xl font-bold"> Password </h2>
                        <div className="w-full flex flex-row justify-center relative mb-6">
                            <Input
                                type={showPassword ? "text" : "password"}
                                id="password"
                                className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-8 focus-visible:ring-0 focus-visible:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                                aria-label="text"
                                value={passw}
                                onChange={(e:ChangeEvent<HTMLInputElement>) => {
                                    setPassword(e.target.value);
                                    if (passw.length < 6) {
                                        setTooShort(true);
                                    } else {
                                        setTooShort(false);
                                    }
                                }}
                            />
                            <Button
                                type="button"
                                onClick={() : void => setShowPassword(!showPassword)}
                                className="absolute top-1/2 right-32 transform -translate-y-1/2 bg-transparent focus-visible:!border-black focus-visible::border focus-visible:!border-8 hover:bg-transparent shadow-none p-2"
                                aria-label={showPassword ? "Hide password" : "Show password"}
                                tabIndex={0}
                            >
                                {showPassword ? <FaEyeSlash className="text-gray-500 scale-x-[-1] text-2xl"/> :
                                    <FaEye className="text-gray-500 scale-x-[-1] text-2xl"/>}
                            </Button>
                        </div>
                        {passwTooShort &&
                                <p className={"text-xl text-red-600 font-bold"}> Your password must be longer than 6 characters</p>
                        }
                            <Button
                                className="text-xl font-bold duration-200 hover:bg-white hover:text-black focus-visible:border-black focus-visible:bg-white  focus-visible:text-black focus-visible:ring-0 focus-visible:border-8 ring-0"
                                onClick={handleDataUpdate}
                                disabled={passw.length < 6}
                            >
                                Save & Update Profile
                            </Button>
                    </div>
                    userData
                    userData

                    <div
                        className="rounded-3xl w-full lg:w-1/2 bg-slate-700 flex flex-col justify-center items-center"
                    >
                        <h3 className="p-4 text-white text-4xl font-bold mb-20"> Linked Accounts</h3>
                        <p className="text-white text-xl text-wrap font-bold p-4"> You can manage here all your external
                            accounts linked to AREA</p>
                        <p className="text-white text-xl text-wrap pb-4 mt-2 mb-12"> Scroll through the supported
                            services </p>

                        <ScrollArea className="h-80 w-72 lg:w-96 bg-white rounded-md border opacity-90 mb-4">
                            <h4 className="mb-4 text-3xl leading-none text-black font-bold m-4"> Services </h4>
                            <div className="flex flex-col items-center p-2 gap-2">
                                {tags.map((tag: string) => (
                                    <div key={tag}>
                                        <div className="flex flex-row items-center justify-between">
                                            <div className="flex items-center">
                                                <div className="my-3">
                                                    <ServiceIcon className="text-2xl" tag={tag}/>
                                                </div>
                                                <div className="mx-4 lg:mx-8 text-2xl text-black font-semibold">
                                                    {tag.charAt(0).toUpperCase() + tag.slice(1)}
                                                </div>
                                            </div>
                                            <div className={"ml-auto"}>
                                                {providers.includes(tag) ?
                                                    <Button
                                                        className="my-2 mr-6 lg:mr-2 font-bold bg-red-600 w-24 focus-visible:border-8 focus-visible:border-black focus-visible:ring-0"
                                                        onClick={() => {
                                                            handleDisconnectProvider(tag);
                                                            setTimeout(() => {
                                                                reloadPage()
                                                            }, 500);
                                                        }}
                                                    >
                                                        Unlink
                                                    </Button> :
                                                    <OauthButton
                                                        arial-label={`${tag}`}
                                                        service={`${tag}`}
                                                        login={false}
                                                        textButton={"Link"}
                                                        className="my-2 mr-6 lg:mr-2 font-bold bg-green-600 w-24 focus-visible:border-8 focus-visible:border-black focus-visible:ring-0"
                                                        ServiceIcon={null}
                                                    />
                                                }
                                            </div>
                                        </div>
                                        <hr className="w-72 h-px bg-black border-0 dark:bg-gray-700"/>
                                    </div>
                                ))}
                            </div>
                        </ScrollArea>
                    </div>
                </div>
            </div>
        </>
    );
}
