'use client';
import { ComboboxDemo } from "@/components/ui/ComboboxDemo";
import { MicroServiceCard } from "@/components/ui/microserviceCard";
import * as React from "react";
import { useEffect, useState, useMemo } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { create } from "@/api/createArea";

const getColorForService = (refName: string) => {
    const colors: { [key: string]: string; } = {
        dt: "green",
        hf: "orange",
        github: "gray"
    };
    return colors[refName] || "blue";
};

const processResponseData = (data) => {
    const services = data.map(item => ({
        value: item.ref_name,
        label: item.name
    }));

    const microservices = data.map(item => ({
        key: item.ref_name,
        color: getColorForService(item.ref_name),
        actions: item.microservices?.map((ms, index) => ({
            id: index,
            title: ms.ref_name,
            description: ms.name,
            type: ms.type,
            ingredients: ms.ingredients
        }))
    }));

    return { services, microservices };
};

function renderIngredientsInput(ingredients: Map<string, string> | undefined, values: string[], setValues: React.Dispatch<React.SetStateAction<string[]>>) {
    if (ingredients === undefined) {
        return <div></div>
    }

    return (
        <>
            {Object.keys(ingredients).map((ingredient, index) => (
                <div key={index} className="flex flex-col justify-center items-center">
                    <p className="p-2 left-0 text-2xl text-white">  {ingredient.charAt(0).toUpperCase() + ingredient.slice(1)} </p>
                    <Input
                        type="text"
                        id={`text-${index}`}
                        className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                        aria-label="text"
                        value={values[index] || ''}
                        onChange={(e) => {
                            const newValues = [...values];
                            newValues[index] = e.target.value;
                            setValues(newValues);
                        }}
                        required
                    />
                </div>
            ))}
        </>
    );
}

const filterAreaByType = (services: AreaServices[], type: string) => {
    return services.filter((service) => {
        return service.microservices.find((micro) => {
            return micro.type == type
        }) != undefined
    }).map((service) => {
        return {
            name: service.name,
            ref_name: service.ref_name,
            microservices: service.microservices.filter((micro) => {
                return micro.type == type
            })
        }
    })
}

const filterServiceByRefName = (services: AreaServices[], refName: string): AreaServices | undefined => {
    return services.find((service) => service.ref_name === refName)
}

export default function CreatePage({ services, uid }: { services: AreaServices[], uid: number }) {
    const actions: AreaServices[] = useMemo(() => {
        return filterAreaByType(services, "action")
    }, [services])

    const [actionName, setActionName] = React.useState("");
    const [microActionName, setMicroActionName] = React.useState("");

    const actionServiceChosen = useMemo(() => {
        return filterServiceByRefName(actions, actionName)
    }, [actions, actionName, setActionName])

    const reactions: AreaServices[] = useMemo(() => {
        return filterAreaByType(services, "reaction")
    }, [services])

    const [reactionName, setReactionName] = React.useState(""); // Later it will be an array of strings
    const [microReactionName, setMicroReactionName] = React.useState("");

    const reactionServiceChosen = useMemo(() => {
        // Loop here with an array of reactionServiceChosen
        return filterServiceByRefName(reactions, reactionName)
    }, [reactions, reactionName, setReactionName])

    const [ingredientValuesActions, setIngredientValuesActions] = useState<string[]>([]);
    const [ingredientValuesReactions, setIngredientValuesReactions] = useState<string[]>([]);

    useEffect(() => {
        setMicroActionName("")
    }, [actionName])

    useEffect(() => {
        setMicroReactionName("")
    }, [reactionName])

    const renderMicroservices = (service: AreaServices | undefined, setMicroservice: (microName: string) => void) => {
        if (service === undefined) {
            return <div></div>
        }
        return (
            <div className="flex flex-wrap py-4 justify-center items-center">
                {service.microservices.map((micro, idx) =>
                    <div key={`${micro.name}-${micro.ref_name}`} className="flex flex-row">
                        <MicroServiceCard
                            setMicroservice={() => { setMicroservice(micro.ref_name) }}
                            microServicesColor={"red"}
                            title={micro.name}
                            description={`Service ${micro.ref_name}`}
                            microserviceName={micro.ref_name}
                        />
                    </div>
                )}
            </div>
        )
    }


    const handleSubmit = () => {
    //     if (microActionName === -1 || microReactionName === -1) {
    //         console.error("Both action and reaction must be selected");
    //         return;
    //     }

    //     const actionMicroservice = microservices
    //         .find(ms => ms.key === actionName)
    //         ?.actions[microActionName];
    //     const reactionMicroservice = microservices
    //         .find(ms => ms.key === reactionName)
    //         ?.actions[microReactionName];

    //     if (!actionMicroservice || !reactionMicroservice) {
    //         console.error("Microservices not found");
    //         return;
    //     }

    //     const actionIngredients = Object.keys(actionMicroservice.ingredients || {});
    //     const reactionIngredients = Object.keys(reactionMicroservice.ingredients || {});

    //     const payload = {
    //         user_id: uid,
    //         action: {
    //             service: actionName,
    //             microservice: actionMicroservice.title,
    //             ingredients: actionIngredients.reduce((acc, key, index) => {
    //                 acc[key] = ingredientValuesActions[index];
    //                 return acc;
    //             }, {})
    //         },
    //         reaction: {
    //             service: reactionName,
    //             microservice: reactionMicroservice.title,
    //             ingredients: reactionIngredients.reduce((acc, key, index) => {
    //                 acc[key] = ingredientValuesReactions[index];
    //                 return acc;
    //             }, {})
    //         }
    //     };

    //     // console.log(payload);
    //     // create(payload)
    };

    return (
        <div className="my-16 bg-white h-full w-full flex flex-col justify-center items-center p-8">
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                <h1 className="my-2 text-blue-500">ACTION</h1>
                <ComboboxDemo services={actions} serviceName={actionName} setValue={setActionName} />
                {actionName === "" && (
                    <h1 className="p-6 text-blue-500 text-5xl">Veuillez sélectionner une action</h1>
                )}
                {microActionName === "" ? (
                    renderMicroservices(
                        actionServiceChosen,
                        (microName) => { setMicroActionName(microName) }
                    )
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {actionServiceChosen && actionServiceChosen.microservices.find((ms) => ms.ref_name === microActionName)?.name || microActionName}
                        </div>
                        <form>
                            {actionServiceChosen && renderIngredientsInput(
                                actionServiceChosen.microservices.find((ms) => ms.ref_name === microActionName)?.ingredients,
                                ingredientValuesActions,
                                setIngredientValuesActions
                            )}
                        </form>
                    </>
                )}
            </div>
            <hr className="h-32 w-4 bg-gray-300" />
            <div // Later for multiple reactions this will be a loop
                className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                <h1 className="my-2 text-red-500">REACTION</h1>
                <ComboboxDemo services={reactions} serviceName={reactionName} setValue={setReactionName} />
                {microReactionName === "" ? (
                    renderMicroservices(
                        reactionServiceChosen,
                        (microName) => { setMicroReactionName(microName) }
                    )
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {reactionServiceChosen && reactionServiceChosen.microservices.find((ms) => ms.ref_name === microReactionName)?.name || microReactionName}
                        </div>
                        <form>
                            {reactionServiceChosen && renderIngredientsInput(
                                reactionServiceChosen.microservices.find((ms) => ms.ref_name === microReactionName)?.ingredients,
                                ingredientValuesReactions,
                                setIngredientValuesReactions
                            )}
                        </form>
                    </>
                )}
            </div>
            <Button
                onClick={handleSubmit}
                className="mt-8 px-6 py-3 bg-green-500 text-white rounded-lg text-3xl font-bold"
                disabled={microActionName === "" || microReactionName === ""}
            >
                Create AREA
            </Button>
        </div>
    );
}
