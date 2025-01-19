'use client';
import { AreaCard } from "@/components/ui/services/AreaCard"
import {Input} from "@/components/ui/utils/Input";
import {FaSearch} from 'react-icons/fa';
import {useState} from "react";
import Image from "next/image";

export default function Explore() {
    const [searchField, setSearchField] = useState("");

    // a changer par un call au backend to get every default area
    let areas = [
        {id: 2, reactionCategory: "Github", areaColor: "green", areaCategory: "Clock", areaTitle: "Change la description et/ou le titre d'un repo github tout les X temps",
            ingredients: {owner: "moi", repo: "area", name: "nom", description: "desc"},
        microAction: "dateTimeTrigger", microReaction: "updateRepo"},
       ];

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchField(e.target.value);
    };

    areas = areas.filter(area => {
        return area.areaTitle.toLowerCase().includes(searchField.toLowerCase());
    })

    return (
        <>
            <div className="pt-36 flex flex-col sm:flex-row justify-center items-center p-32">
                <Image
                    className="opacity-40"
                    src="/bg.png"
                    alt="Explore AREAs"
                    width="350"
                    height="350"
                />
                <div
                    className="hidden flex-col sm:block sm:flex-col justify-center items-center p-4 text-2xl font-bold"
                >
                    <h1 className="mx-16 my-8 py-2">
                        Discover a world of ready-to-use automation
                    </h1>
                    <h2 className="mx-20 flex flex-row items-center">
                        Explore our
                        <p className="mx-2 font-bold text-green-500 animate-pulse ease-in-out duration-1000">
                            Pre-defined AREAs
                        </p>
                    </h2>
                    <h3 className="mx-20 flex flex-row items-center">
                        From
                        <p className="mx-1 font-bold text-blue-500 animate-bounce ease-in-out">
                            ACTION
                        </p>
                        to
                        <p className="mx-1 font-bold text-red-500 animate-pulse">
                            REACTION
                        </p>
                        everything is set up for you
                    </h3>
                    <p className="mx-20 mt-4 text-xl font-normal text-gray-600">
                        Choose, customize, and activate with just a few clicks
                    </p>
                </div>
            </div>

            <div className="bg-gray-50 flex flex-col items-center justify-center">
                <h6 className="my-8 text-5xl text-black font-extrabold"> Explore </h6>
                <div
                    className="sm:1/3 flex focus-within:border-black flex-row items-center justify-center rounded-2xl my-8 font-extrabold focus-visible:border-black border-4 p-4 h-16 bg-slate-300 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                >
                    <FaSearch className="text-gray-400 text-3xl"/>
                    <Input
                        placeholder="Cherchez des Areas ici"
                        id="Search"
                        className="!text-2xl border-none shadow-none focus-visible:border-none focus-visible:!ring-0 focus-visible:shadow-none focus-visible:outline-none"
                        aria-label="Search for area"
                        onChange={handleChange}
                    />
                </div>
                <div className="mx-64 h-1/2 w-3/4 flex flex-wrap items-center justify-center">
                    {areas.map(area => (
                        <div key={area.id} className="">
                            <AreaCard
                                areaColor={area.areaColor} category={area.areaCategory} areaTitle={area.areaTitle}
                                action={area.areaCategory} reaction={area.reactionCategory}
                            />
                        </div>
                    ))}
                </div>
                <div className="p-16">

                </div>
            </div>
        </>
    )
}
