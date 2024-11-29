'use client';
import { CustomCard } from "@/components/ui/customCard"
import {Input} from "@/components/ui/input";
import {FaSearch} from 'react-icons/fa';

export default function Page() {

    interface Area {
        areaCategory: string;
        areaTitle: string;
        actionImage: string;
        reactionImage: string;
    }

    const areas = [
        {id: 0, areaColor: "red", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/6/icons/monochrome_regular.webp"},
        {id: 1, areaColor: "blue", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/51464135/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 2, areaColor: "green", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 3, areaColor: "purple", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/51464135/icons/monochrome_regular.webp"},
        {id: 4, areaColor: "red", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 5, areaColor: "red", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/6/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
    ];

    return (
        <>
            <div className="flex flex-col items-center justify-center">
                <h1 className="my-8 text-5xl text-black font-extrabold"> Explore </h1>
                <div className="w-1/3 flex focus-within:border-black flex-row items-center justify-center rounded-2xl my-8 font-extrabold focus:border-black border-4 p-4 h-16 bg-slate-300 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60">
                    <FaSearch className="text-gray-400 text-3xl" />
                    <Input
                        placeholder="Cherchez des Areas ici"
                        id="Search"
                        className="!text-2xl border-none shadow-none focus:border-none focus:!ring-0 focus:shadow-none focus:outline-none"
                        aria-label="Search for area"
                    />
                </div>
                <div className="mx-64 h-1/2 w-3/4 flex flex-wrap items-center justify-center">
                    {areas.map(area => (
                        <div key={area.id} className="">
                            <CustomCard areaColor={area.areaColor} category={area.areaCategory} areaTitle={area.areaTitle} action={area.actionImage} reaction ={area.reactionImage} />
                        </div>
                    ))}
                </div>
            </div>
        </>
    )
}
