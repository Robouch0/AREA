'use client';
import {Card, CardHeader, CardTitle, CardDescription} from "@/components/ui/services/Card";
import {Button} from "@/components/ui/utils/Button";
// import { useRouter } from 'next/navigation';
import {ServiceIcon} from "@/components/ui/services/ServiceIcon";
import {getColorForService} from "@/lib/utils";
import {AreaServices} from "@/api/types/areaStatus";

export function MyAreaCard({action, reaction, areaID}: {
    action: AreaServices,
    reaction: AreaServices|undefined,
    areaID: number,
}) {
    console.log(areaID)
    if (reaction === undefined) {
        return <h1 className="text-black text-7xl mt-20"> ERROR with the Reaction </h1>
    }
    return (
        <>
            <Button
                className="hover:bg-transparent shadow-none bg-transparent flex flex-row my-2 w-96 h-96 focus-visible:border-slate-500 focus-visible:border-8 rounded-3xl focus-visible:outline-none"
            >
                <Card
                    className={"font-black text-3xl w-full h-full border-none hover:opacity-75 flex flex-col justify-between"}
                    style={{backgroundColor: getColorForService(action.ref_name)}}
                >
                    <div className="flex flex-col">
                        <div className="flex flex-row py-2 px-2 my-2 items-center">

                            <ServiceIcon className="text-2xl text-black mx-2" tag={action.ref_name}/>

                            <ServiceIcon className="text-3xl text-black" tag={reaction.ref_name}/>
                        </div>
                        <CardHeader className="text-wrap mt-2">
                            <CardTitle className="text-blue-600 my-2 !text-2xl break-words mb-6"> {action.microservices?.at(0)?.name} </CardTitle>
                            <CardTitle className="text-red-600 my-2 !text-2xl break-words"> {reaction.microservices.at(0)?.name} </CardTitle>

                        </CardHeader>
                    </div>
                    <div className="flex flex-row py-3 mx-4">
                        <ServiceIcon tag={action.ref_name}/>

                        <CardDescription className="mx-3 text-xl text-black">
                            {action.name}
                        </CardDescription>
                    </div>
                </Card>

            </Button>
        </>
    );
}
