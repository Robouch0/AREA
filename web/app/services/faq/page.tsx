"use client";
import {FaCircleQuestion} from "react-icons/fa6";
import {Button} from "@/components/ui/button";

export default function Faq() {
    return (
        <>
            <p className="mt-32 text-center font-bold text-3xl">Frequently Asked Questions:</p>

            <div className="mt-16 flex flex-col items-center">
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
                        <div className="flex flex-row justify-center items-center">
                            <p className={"font-medium text-black text-xl"}>
                                You can also download the pdf of the documentation
                            </p>
                            <Button
                                onClick={() => window.open(`/bg.png`)}
                            >
                                Here
                            </Button>
                        </div>
                    </div>
                </div>

                <div className="w-full max-w-2xl">
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
                </div>
            </div>
        </>
    );
}
