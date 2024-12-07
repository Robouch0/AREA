'use client';
import { AreaCard } from "@/components/ui/areaCard"
import {Input} from "@/components/ui/input";
import {FaSearch} from 'react-icons/fa';
import {useState} from "react";
import Image from "next/image";

export default function Explore() {
    const [searchField, setSearchField] = useState("");

    // exemple mockup
    // a changer par un call au backend to get every default area
    let areas = [
        {id: 0, areaColor: "red", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/6/icons/monochrome_regular.webp"},
        {id: 1, areaColor: "blue", areaCategory: "Youtube", areaTitle: "Nouvel vidéo d'inoxtag, play quoicoubébou des montagnes", actionImage: "https://assets.ifttt.com/images/channels/51464135/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 2, areaColor: "green", areaCategory: "Instagram", areaTitle: "nouvelle recette sur marmiton, tweetez le", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 3, areaColor: "purple", areaCategory: "Instagram", areaTitle: "TOut les vendredis de chaque semaine récupérer mon top 10 spotify", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/51464135/icons/monochrome_regular.webp"},
        {id: 4, areaColor: "red", areaCategory: "Instagram", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 5, areaColor: "red", areaCategory: "Outlook", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/6/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
        {id: 6, areaColor: "orange", areaCategory: "Tweeter", areaTitle: "If Every hour at 00 minutes past the hour, then Send me an email at hugo.duchemin.r@gmail.com", actionImage: "https://assets.ifttt.com/images/channels/6/icons/monochrome_regular.webp", reactionImage: "https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"},
    ];

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchField(e.target.value);
    };

    areas = areas.filter(area => {
        return area.areaTitle.toLowerCase().includes(searchField.toLowerCase());
    })

    return (
        <>
            <div className="bg-white flex flex-col sm:flex-row justify-center items-center p-32">
                <Image
                    className="opacity-40"
                    src="/bg.png"
                    alt="bg Social Networks"
                    width="350"
                    height="350"
                />
                <div className="hidden flex-col sm:block sm:flex-col justify-center items-center p-4 text-2xl font-bold">
                    <h1 className="mx-16 my-8 py-2">
                        Le monde de l&#39;automatisation s&#39;ouvre pour vous
                    </h1>
                    <h2 className="mx-20 flex flex-row">
                        Il est l&#39;heure de passer à l&#39;
                        <p className="mx-1 my-1 font-bold text-blue-500 animate-bounce ease-in-out">
                            ACTION
                        </p>
                    </h2>
                    <h3 className="mx-20 flex flex-row">
                        Nous avons hâte de voir vos
                        <p className="mx-2 font-bold text-red-500 animate-pulse">
                            REACTIONS
                        </p>
                    </h3>
                </div>
            </div>
            <div className="flex flex-col items-center justify-center bg-white">
                <h6 className="my-8 text-5xl text-black font-extrabold"> Explore </h6>
                <div
                    className="sm:1/3 flex focus-within:border-black flex-row items-center justify-center rounded-2xl my-8 font-extrabold focus:border-black border-4 p-4 h-16 bg-slate-300 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60">
                    <FaSearch className="text-gray-400 text-3xl" />
                    <Input
                        placeholder="Cherchez des Areas ici"
                        id="Search"
                        className="!text-2xl border-none shadow-none focus:border-none focus:!ring-0 focus:shadow-none focus:outline-none"
                        aria-label="Search for area"
                        onChange={handleChange}
                    />
                </div>
                <div className="mx-64 h-1/2 w-3/4 flex flex-wrap items-center justify-center">
                    {areas.map(area => (
                        <div key={area.id} className="">
                            <AreaCard areaColor={area.areaColor} category={area.areaCategory} areaTitle={area.areaTitle} action={area.actionImage} reaction ={area.reactionImage} />
                        </div>
                    ))}
                </div>
                <div className="p-16">

                </div>
            </div>
        </>
    )
}
