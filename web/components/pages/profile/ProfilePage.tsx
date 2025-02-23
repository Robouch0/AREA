"use client";

import Image from "next/image";
import {Input} from "@/components/ui/utils/thirdPartyComponents/shadcn/Input";
import {Button} from "@/components/ui/utils/thirdPartyComponents/shadcn/Button";
import {ScrollArea} from "@/components/ui/utils/thirdPartyComponents/shadcn/Scroll-area"
import {ChangeEvent, useState} from "react";
import {ServiceIcon} from "@/components/ui/utils/ServiceIcon";
import {userInfo} from "@/api/getUserInfos";
import {OauthButton} from "@/components/ui/services/oauth/OauthButton";
import {unlinkOauthProvider} from "@/api/unlinkOauth";
import {updateUserInfos} from "@/api/updateUserInfos";
import {useToast} from "@/hooks/use-toast";
import {VideoTutorialPopUp} from "@/components/ui/utils/VideoTutorialPopUp";


export default function ProfilePage({id, email, first_name, last_name, password, usersProviders, possibleProviders}: userInfo) {
    const [firstName, setFirstName] = useState(first_name);
    const [lastName, setLastName] = useState(last_name);
    const [passw] = useState(password);
    const { toast } = useToast()

    possibleProviders.sort((a, b) => a.localeCompare(b))

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
            })
        }
        ).catch(() => {
            toast({
                title: "Update failed",
                description: "Your new datas have not been updated on our server.",
                variant: 'destructive',
                duration: 2500,
            })
        }
        )
    }

    const reloadPage = () => {
        window.location.reload()
    }

    if (id == -1) {
        toast({
            title: "The area server is down",
            description: "Cannot get your profile datas for the moment",
            variant: 'destructive',
            duration: 2500,
        })
        return <> </>;
    }

    return (
        <>
            <div className="pt-10 w-full h-full flex flex-col bg-white justify-center items-center my-16">
                <h1 className="text-6xl font-bold p-8"> Account Settings </h1>
                <hr className="w-1/3  my-8 bg-slate-800 opacity-20 h-1 border-0 dark:bg-gray-700" />
                <div
                    className="w-4/5 h-full bg-slate-800 rounded-3xl flex flex-col lg:flex-row p-8 m-2 justify-between"
                >
                    <div
                        className="p-6 rounded-3xl w-full lg:w-1/2 bg-slate-600 h-full flex flex-col justify-center items-center gap-4"
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
                            onChange={(e: ChangeEvent<HTMLInputElement>) => {
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
                            onChange={(e: ChangeEvent<HTMLInputElement>) => {
                                setLastName(e.target.value);
                            }}
                        />
                        <div className={"flex flex-row text-white justify-between items-center w-full"}>
                            <div
                                className={"flex-1"}
                                aria-label={"Tutorial video, how to update your private datas"}
                            >
                                <VideoTutorialPopUp description={"How to update your private datas"} videoPath={"/tutoUpdateProfileDatas.mp4"}/>
                            </div>
                            <Button
                                className="flex justify-center text-xl font-bold duration-200 hover:bg-white hover:text-black focus-visible:border-black focus-visible:bg-white  focus-visible:text-black focus-visible:ring-0 focus-visible:border-8 ring-0"
                                aria-label={"save & update your personnal data"}
                                onClick={handleDataUpdate}
                            >
                                Save & Update Profile
                            </Button>
                            <div className={"flex-1"}></div>
                        </div>
                    </div>


                    <div
                        className="rounded-3xl w-full lg:w-1/2  flex flex-col justify-center items-center"
                    >
                        <div className={"flex flex-row pb-4 text-white justify-between items-center w-full"}>
                            <div className={"flex-1"}></div>
                            <h3 className="flex justify-center text-white text-4xl font-bold"> Linked Accounts</h3>
                            <div className={"flex-1"}>
                                <VideoTutorialPopUp description={"How to link another account to Area"} videoPath={"/tutoLinkOauthAccount.mp4"}/>
                            </div>
                        </div>
                        <p className="text-white text-xl text-wrap font-bold p-4"> Manage here all your external
                            accounts linked to AREA</p>
                        <p className="text-white text-xl text-wrap pb-4 mt-2 mb-12"> Scroll through the supported
                            services </p>

                        <ScrollArea className="h-80 w-72 lg:w-96 bg-white rounded-md border opacity-90 mb-4">
                            <h4 className="mb-4 text-3xl leading-none text-black font-bold m-4"> Services </h4>
                            <div className="flex flex-col items-center p-2 gap-2">
                                {possibleProviders.map((tag: string) => (
                                    <div key={tag}>
                                        <div className="flex flex-row items-center justify-between">
                                            <div className="flex items-center">
                                                <div className="my-3">
                                                    <ServiceIcon className="text-2xl" tag={tag} />
                                                </div>
                                                <div className="mx-4 lg:mx-8 text-2xl text-black font-semibold">
                                                    {tag.charAt(0).toUpperCase() + tag.slice(1)}
                                                </div>
                                            </div>
                                            <div className={"ml-auto"}>
                                                {usersProviders !== null && usersProviders.includes(tag) ?
                                                    <Button
                                                        className="my-2 mr-6 lg:mr-2 text-black text-xl font-bold bg-red-600 w-24 focus-visible:border-8 focus-visible:border-black focus-visible:ring-0"
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
                                                        location="profile"
                                                        textButton={"Link"}
                                                        className="my-2 mr-6 lg:mr-2 font-bold bg-green-600 w-24 focus-visible:border-8 focus-visible:border-black focus-visible:ring-0"
                                                        ServiceIcon={null}
                                                    />
                                                }
                                            </div>
                                        </div>
                                        <hr className="w-72 h-px bg-black border-0 dark:bg-gray-700" />
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
