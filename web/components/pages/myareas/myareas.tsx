'use client';
import { AreaCard } from "@/components/ui/services/AreaCard"
import {Input} from "@/components/ui/utils/Input";
import {FaSearch} from 'react-icons/fa';
import {useState} from "react";
import Image from "next/image";
import {AreaServicesWithId} from "@/api/types/areaStatus";
import {MyAreaCard} from "@/components/ui/services/MyAreaCard";

export default function Myareas({ userAreas }: { userAreas: AreaServicesWithId[] }) {
    const [searchField, setSearchField] = useState("");

    const filteredAreas = userAreas.filter(area => {
        const actionNameMatch = area.Action.name.toLowerCase().includes(searchField.toLowerCase());
        const reactionNameMatch = area.Reactions?.[0]?.name.toLowerCase().includes(searchField.toLowerCase());
        const areaName = area.Action.microservices?.at(0)?.name.toLowerCase().includes(searchField.toLowerCase());
        return actionNameMatch || reactionNameMatch || areaName;
    })
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setSearchField(e.target.value);
    };

    return (
        <>
            <div className="pt-36 flex flex-col sm:flex-row justify-center items-center p-32">
                <Image
                    className="opacity-100"
                    src="/puzzle.jpg"
                    alt="bg Social Networks"
                    width="330"
                    height="330"
                />
                <div className="hidden flex-col sm:block sm:flex-col justify-center items-center p-4 text-2xl font-bold">
                    <h1 className="mx-16 my-8 py-2">
                        Here are all the Areas you created for your account
                    </h1>
                </div>
            </div>
            <div className="bg-gray-50 flex flex-col items-center justify-center">
                <h6 className="my-8 text-5xl text-black font-extrabold"> My Areas </h6>
                <div
                    className="sm:1/3 flex focus-within:border-black flex-row items-center justify-center rounded-2xl my-8 font-extrabold focus-visible:border-black border-4 p-4 h-16 bg-slate-300 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60">
                    <FaSearch className="text-gray-400 text-3xl" />
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
                            <MyAreaCard action={area.Action} reaction={area.Reactions?.at(0)} areaID={parseInt(area.ID)}> </MyAreaCard>
                            {/*<AreaCard areaColor={area.areaColor} category={area.areaCategory} areaTitle={area.areaTitle} action={area.areaCategory} reaction ={area.reactionCategory} />*/}
                        </div>
                    ))}
                </div>
                <div className="p-16">

                </div>
            </div>
        </>
    )
}
