'use client';
import { ComboboxDemo } from "@/components/ui/ComboboxDemo";
import { MicroServiceCard } from "@/components/ui/microserviceCard";
import * as React from "react";
import {useEffect} from "react";
import {testCreateModifyRepo} from "@/api/createArea";

export type PossibleType = string | number | { [key: string]: string } | { [key: string]: PossibleType };


function checkForUrlParams(): Record<string, string>  | undefined{
    const params: Record<string, string> = {};
    const url = window.location.href;
    const queryString = url.split('?')[1];

    if (queryString) {
        const pairs = queryString.split('&');
        pairs.forEach(pair => {
            const [key, value] = pair.split('=');
            if (key && value) {
                params[decodeURIComponent(key)] = decodeURIComponent(value);
            }
        });
    } else {
        return undefined;
    }

    console.log(params);
    return params;
}


export default function Create() {
    const params = checkForUrlParams();

    useEffect(() => {
        switch(params?.reaction) {
            case "updateRepo" :
                testCreateModifyRepo();
                break;
            default:
                console.log("no area for this")
        }
         testCreateModifyRepo();
    }, [params?.reaction]);

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

    const [actionValue, setValue] = React.useState("")
    const [reactionValue, setReactionValue] = React.useState("")
    const [actionId, setActionId] = React.useState(-1)
    const [reactionId, setReactionId] = React.useState(-1)

    useEffect(() => {
        setActionId(-1);
    }, [actionValue]);

    useEffect(() => {
        setReactionId(-1);
    }, [reactionValue]);


    return (
        <div className="my-16 bg-white h-full w-full flex flex-col justify-center items-center p-8">
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl  flex flex-col justify-start items-center">
                <h1 className="my-2 text-blue-500">
                    ACTION
                </h1>
                <ComboboxDemo services={services} value={actionValue} setValue={setValue}/>
                {(actionValue != "" && actionId == -1) ? <h1 className="p-6 text-blue-500 text-5xl">Veuillez sélectionner une action </h1> : <div> </div>}
                { actionId == -1 ?
                    <div className="flex flex-wrap py-4 justify-center items-center">
                        {microservices.find((microservice) => microservice.key === actionValue)?.actions.map((action) => (
                          <div key={action.id}
                              className="flex flex-row">
                              <MicroServiceCard
                                  id={action.id}
                                  setId={setActionId}
                                  microServicesColor={microservices.find((ms) => ms.key === actionValue)?.color}
                                  title={action.title}
                                  description={action.description}
                              />
                          </div>
                        ))}
                    </div>
                    :
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === actionValue)?.actions?.at(actionId)?.title}
                        </div>
                        <div className="text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === actionValue)?.actions?.at(actionId)?.description}
                        </div>
                    </>
                }
            </div>
            <hr className="h-32 w-4 bg-gray-300"/>
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl  flex flex-col justify-start items-center">
                <h1 className="my-2 text-red-500">
                    REACTION
                </h1>
                <ComboboxDemo services={services} value={reactionValue} setValue={setReactionValue}/>
                {(reactionValue != "" && reactionId == -1) ? <h1 className="p-6 text-red-500 text-5xl">Veuillez sélectionner une réaction </h1> : <div> </div>}
                { reactionId == -1 ?
                    <div className="flex flex-wrap py-4 justify-center items-center">
                        {microservices.find((microservice) => microservice.key === reactionValue)?.actions.map((reaction) => (
                            <div key={reaction.id}
                                 className="flex flex-row">
                                <MicroServiceCard
                                    id={reaction.id}
                                    setId={setReactionId}
                                    microServicesColor={microservices.find((ms) => ms.key === reactionValue)?.color}
                                    title={reaction.title}
                                    description={reaction.description}
                                />
                            </div>
                        ))}
                    </div>
                    :
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === reactionValue)?.actions?.at(reactionId)?.title}
                        </div>
                        <div className="text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === reactionValue)?.actions?.at(reactionId)?.description}
                        </div>
                    </>
                }
            </div>
        </div>
    );
}
