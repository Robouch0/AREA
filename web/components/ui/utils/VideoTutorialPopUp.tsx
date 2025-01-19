import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip";
import * as React from "react";
import {FaQuestionCircle} from "react-icons/fa";

interface VideoTutorialPopUpProps {
    description: string;
    videoPath: string;
}

export function VideoTutorialPopUp({description, videoPath} : VideoTutorialPopUpProps) {
    return (
        <TooltipProvider>
            <Tooltip>
                <TooltipTrigger>
                    <div className={"flex flex-row items-center justify-center"}>
                        <FaQuestionCircle className={"mx-3 text-2xl"}></FaQuestionCircle>
                    </div>
                </TooltipTrigger>
                <TooltipContent>
                    <div className={"flex flex-col justify-center items-center bg-slate-400"}>
                        <p className={"font-bold text-black p-2 text-3xl"}> {description} </p>
                        <iframe
                            width={844}
                            height={475}
                            allow="autoplay"
                            src={videoPath}
                            title="tutorial"
                        >
                        </iframe>
                    </div>
                </TooltipContent>
            </Tooltip>
        </TooltipProvider>
    )
}
