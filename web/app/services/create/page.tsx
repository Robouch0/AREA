'use client';
import { ComboboxDemo } from "@/components/ui/ComboboxDemo";
import { MicroServiceCard } from "@/components/ui/microserviceCard";
import * as React from "react";
import {useEffect} from "react";

export default function Create() {

    const services = [
        {
            value: "instagram",
            label: "Instagram",
        },
        {
            value: "youtube",
            label: "YouTube",
        },
        {
            value: "gmail",
            label: "Gmail",
        },
        {
            value: "github",
            label: "Github",
        },
    ]

    const microservices = [
        {
            key: "instagram",
            color: "blue",
            actions: [
                {
                    id: 0,
                    title: "Un like sur un de vos post",
                    description: "Ce déclencheur s'active lorsqu'un de vos post se fait like"
                },
            ]
        },
        {
            key: "youtube",
            color: "red",
            actions: [
                {
                    id: 0,
                    title: "Un like sur une de vos vidéos",
                    description: "Ce déclencheur s'active lorsqu'une de vos vidéos s'est fait liké"
                },
                {
                    id: 1,
                    title: "Un mec s'est abonné à vous",
                    description: "Ce déclencheur s'active lorsqu'une personne s'est abonné à votre chaîne"
                },
                {
                    id: 2,
                    title: "Un mec s'est abonné à vous",
                    description: "Ce déclencheur s'active lorsqu'une personne s'est abonné à votre chaîne"
                },
                {
                    id: 3,
                    title: "Un like sur une de vos vidéos",
                    description: "Ce déclencheur s'active lorsqu'une de vos vidéos s'est fait liké"
                },
                {
                    id: 4,
                    title: "Un mec s'est abonné à vous",
                    description: "Ce déclencheur s'active lorsqu'une personne s'est abonné à votre chaîne"
                },
                {
                    id: 5,
                    title: "A ça mon pote, jvais tous les faire",
                    description: "Un peu de haine ici il buvait tout son RSA, il s'appellait COCO à la tournée du bar"
                },
                {
                    id: 6,
                    title: "Un mec s'est abonné à vous",
                    description: "Ce déclencheur s'active lorsqu'une personne s'est abonné à votre chaîne"
                },
            ]
        },
    ]

    const [value, setValue] = React.useState("")
    const [actionId, setActionId] = React.useState(-1)

    useEffect(() => {
        setActionId(-1);
    }, [value]);

    return (
        <div className="my-16 bg-white h-full w-full flex flex-col justify-center items-center p-4">
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl  flex flex-col justify-start items-center">
                <h1 className="my-2 text-blue-500">
                    ACTION
                </h1>
                <ComboboxDemo services={services} value={value} setValue={setValue}/>
                {(value != "" && actionId == -1) ? <h1 className="p-6 text-blue-500 text-5xl">Veuillez sélectionner une action </h1> : <div> </div>}
                { actionId == -1 ?
                    <div className="flex flex-wrap py-4 justify-center items-center">
                        {microservices.find((microservice) => microservice.key === value)?.actions.map((action) => (
                          <div key={action.id}
                              className="flex flex-row">
                              <MicroServiceCard
                                  id={action.id}
                                  setId={setActionId}
                                  microServicesColor={microservices.find((ms) => ms.key === value)?.color}
                                  title={action.title}
                                  description={action.description}
                              />
                          </div>
                        ))}
                    </div>
                    :
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === value)?.actions?.at(actionId)?.title}
                        </div>
                        <div className="text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === value)?.actions?.at(actionId)?.description}
                        </div>
                    </>
                }
            </div>
            <hr className="h-32 w-4 bg-gray-300"/>
            <div className="bg-slate-800 text-6xl font-bold w-1/2 py-4 rounded-3xl h-80 flex flex-col justify-start items-center">
                <h1 className="my-2 text-white">
                    REACTION
                </h1>
            </div>
        </div>
    );
}
