'use client';
import {Input} from "@/components/ui/utils/Input";
import {FaSearch} from 'react-icons/fa';
import {useState} from "react";
import Image from "next/image";
import {AreaServicesWithId} from "@/api/types/areaStatus";
import {MyAreaCard} from "@/components/ui/services/MyAreaCard";

export default function MyAreas({ userAreas }: { userAreas: AreaServicesWithId[] }) {
    const [searchField, setSearchField] = useState("");

    const filteredAreas : AreaServicesWithId[] = userAreas.filter(area => {
        const actionNameMatch : boolean = area.Action.name.toLowerCase().includes(searchField.toLowerCase());
        const reactionNameMatch : boolean = area.Reactions?.[0]?.name.toLowerCase().includes(searchField.toLowerCase());
        const areaName:  boolean|undefined = area.Action.microservices?.at(0)?.name.toLowerCase().includes(searchField.toLowerCase());
        return actionNameMatch || reactionNameMatch || areaName;
    })
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchField(e.target.value);
    };

    return (
        <>
            <div className="container mx-auto px-4 py-12 mt-24">
                <div className="text-xl flex flex-col lg:flex-row items-center justify-center gap-8">
                    <div className="text-xl lg:w-1/3">
                        <Image
                            className="text-xl "
                            src="/puzzle.jpg"
                            alt="Puzzle representing Areas"
                            width={330}
                            height={330}
                        />
                    </div>
                    <div className="text-xl lg:w-2/3 space-y-6">
                        <h1 className="text-3xl font-bold text-gray-800 mb-4">
                            Welcome to Your Areas
                        </h1>
                        <p className="text-xl text-gray-600">
                            Here you can find all the Areas you&#39;ve created for your account. Each Area represents a
                            unique automation you&#39;ve set up.
                        </p>
                        <div className="text-xl space-y-2">
                            <p className="text-xl text-gray-600">
                                At the bottom of each Area card, you&#39;ll find:
                            </p>
                            <ul className="text-xl list-disc list-inside text-gray-600 ml-4">
                                <li>The category of the action type it belongs to</li>
                                <li>A quick visual representation of its utility</li>
                            </ul>
                        </div>
                        <div className="text-xl flex items-center space-x-2">
                            <span className="text-xl text-gray-600">The</span>
                            <span className="font-bold text-blue-600">blue </span>
                            <span className="text-xl text-gray-600">represents the name of the action microservice, while the</span>
                            <span className="font-bold text-red-600">red </span>
                            <span className="text-xl text-gray-600">shows the reaction microservice.</span>
                        </div>
                    </div>
                </div>
            </div>

            <div className="bg-gray-50 flex flex-col items-center justify-center">
                <h6 className="my-8 text-5xl text-black font-extrabold"> My Areas </h6>
                <div
                    className="sm:1/3 flex focus-within:border-black flex-row items-center justify-center rounded-2xl my-8 font-extrabold focus-visible:border-black border-4 p-4 h-16 bg-slate-300 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                >
                    <FaSearch className="text-gray-400 text-3xl"/>
                    <Input
                        placeholder="Search through your areas"
                        id="Search"
                        className="!text-2xl border-none shadow-none focus-visible:border-none focus-visible:!ring-0 focus-visible:shadow-none focus-visible:outline-none"
                        aria-label="Search for area"
                        onChange={handleChange}
                    />
                </div>
                <div className="mx-64 h-1/2 w-3/4 flex flex-wrap items-center justify-center">
                    {filteredAreas.map(area => (
                        <div key={area.ID} className="">
                            <MyAreaCard
                                action={area.Action} reaction={area.Reactions?.at(0)} areaID={parseInt(area.ID)}
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
