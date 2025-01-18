'use client';
import {Card, CardHeader, CardTitle, CardDescription} from "@/components/ui/services/Card";
import {ServiceIcon} from "@/components/ui/services/ServiceIcon";
import {getColorForService} from "@/lib/utils";
import {AreaServices} from "@/api/types/areaStatus";
import {Switch} from "@/components/ui/switch";
import {useState} from "react";
import {activateArea} from "@/api/enableArea";

export function MyAreaCard({action, reactions, areaID, areaActivate}: {
    action: AreaServices,
    reactions: AreaServices[],
    areaID: number,
    areaActivate: boolean,
}) {
    const [switchOn, setSwitchOn] = useState<boolean>(areaActivate)

    const handleSwitchToggle = async () => {
        setSwitchOn(!switchOn)
        console.log(areaID)
        activateArea(areaID, !switchOn).then((res) => {
            console.log(res)
        }).catch((err) => console.log(err));
    };

    const mappedReactServices = reactions?.map((react, index) => {
        console.log(react.name)
        return (
            <div key={index}>
                <p className="text-red-700 my-2 !text-2xl break-words"> {react.microservices?.at(0)?.name} </p>
            </div>
        );
    });

    return (
        <>
            <div
                className="p-4 hover:bg-transparent shadow-none bg-transparent flex flex-row my-2 w-96 h-96 focus-visible:border-slate-500 focus-visible:border-8 rounded-3xl focus-visible:outline-none"
                tabIndex={0}
            >
                <Card
                    className={"font-black text-3xl w-full h-full border-none flex flex-col justify-between overflow-y-auto"}
                    style={{backgroundColor: getColorForService(action.ref_name)}}
                >
                    <div className="flex flex-col">
                        <div className="flex flex-row py-2 px-2 my-2 items-center">

                            <ServiceIcon className="text-2xl text-black mx-2" tag={action.ref_name}/>

                            <ServiceIcon className="text-3xl text-black" tag={reactions?.at(0)?.name}/>
                            <div className="flex items-center ml-auto" >
                                <Switch
                                    className={`transition-colors duration-300 ${switchOn ? '!bg-green-500' : '!bg-red-500'}`}
                                    id={`switch-${areaID}`} onClick={handleSwitchToggle} checked={switchOn}/>
                            </div>
                        </div>
                        <CardHeader className="text-wrap mt-2">
                            <CardTitle
                                className="text-blue-700 my-2 !text-2xl break-words mb-6"
                            > {action.microservices?.at(0)?.name} </CardTitle>
                            {mappedReactServices}

                        </CardHeader>
                    </div>
                    <div className="flex flex-row py-3 mx-4">
                        <ServiceIcon tag={action.ref_name}/>

                        <CardDescription className="mx-3 text-xl text-black">
                            {action.name}
                        </CardDescription>

                    </div>
                </Card>

            </div>
        </>
    );
}
