"use client";
import {FaCircleQuestion} from "react-icons/fa6";
import {Button} from "@/components/ui/utils/thirdPartyComponents/shadcn/Button";
import {FaDownload} from "react-icons/fa";

export default function Faq() {
    return (
        <>
            <p className="mt-32 text-center font-bold text-3xl">Frequently Asked Questions:</p>

            <div className="mt-16 flex flex-col items-center mb-8">
                <div className="w-full max-w-2xl mb-8">
                    <p className="font-bold text-2xl text-black">Is there any tutorial on how to use the website?</p>
                    <div className={"ml-4 mt-4"}>
                        <div className="flex items-center">
                            <p className="font-medium text-black text-xl">On every page, you will find those kinds of
                                symbols:</p>
                            <FaCircleQuestion className="ml-2 text-xl"/>
                        </div>
                        <p className="font-medium text-black text-xl">
                            If you hover your cursor on them, they will provide a short popup video on how to use the
                            features of the current page.
                        </p>
                        <div className="flex flex-row justify-start items-center">
                            <p className={"font-medium text-black text-xl left-0"}>
                                You can also download the pdf of our user guide
                            </p>
                            <Button
                                className={"mx-4 h-8 font-bold flex flex-row items-center justify-center hover:opacity-70 focus-visible:border-4 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 duration-200"}
                                onClick={() => window.open(`/userDocumentation.pdf`)}
                            >
                                Here
                                <FaDownload className={"!text-lg"}/>
                            </Button>
                        </div>
                    </div>
                </div>

                <div className="w-full max-w-2xl mb-8 mt-8">
                    <p className="font-bold text-2xl text-black">Where can I link external accounts to my Area
                        account?</p>
                    <div className={"ml-4 mt-4"}>
                        <p className="font-medium text-black text-xl">You can link accounts through OAuth in your
                            profile.</p>
                        <p className="font-medium text-black text-xl">
                            You will also be asked to do it on the create page if you try to use a microservice from a
                            service that you are not connected to.
                        </p>
                    </div>

                    <div className="mt-16 flex flex-col items-center mb-8">
                        <div className="w-full max-w-2xl mb-8">
                            <p className="font-bold text-2xl text-black"> Do you also have a mobile app ?</p>
                            <div className={"ml-4 mt-4"}>
                                <div className="flex items-center">
                                    <p className="font-medium text-black text-xl"> Yes we do ! But it is not verified
                                        yet on
                                        the app store or play store </p>
                                </div>
                                <div className="flex flex-row justify-start items-center">
                                    <p className={"font-medium text-black text-xl left-0"}>
                                        However you can download the ready to use apk
                                    </p>
                                    <Button
                                        className={"mx-4 h-8 font-bold flex flex-row items-center justify-center hover:opacity-70 focus-visible:border-4 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 duration-200"}
                                        onClick={() => window.open(`/client.apk`)}
                                    >
                                        Here
                                        <FaDownload className={"!text-lg"}/>
                                    </Button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>


            </div>
        </>
    );
}
