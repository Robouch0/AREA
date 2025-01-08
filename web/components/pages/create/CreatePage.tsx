'use client';
import { MicroServiceCard } from "@/components/ui/services/MicroserviceCard";
import * as React from "react";
import { useEffect, useState, useMemo } from "react";
import { Button } from "@/components/ui/utils/Button";
import { create } from "@/api/createArea";
import Form from 'next/form';
import MicroserviceCreateZone from "@/components/ui/services/MicroserviceCreateZone";
import { getColorForService } from "@/lib/utils";
import { AreaServices, AreaMicroservices, Ingredient } from "@/api/types/areaStatus";
import { AreaCreateBody } from "@/api/types/areaCreateBody";
import { getUserTokens } from "@/api/getUserInfos";
import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip";
import { InputFieldComponent} from "@/components/pages/create/InputFieldComponent";
import {useRouter} from "next/navigation";
import {useToast} from "@/hooks/use-toast";

export function renderMicroservices(service: AreaServices | undefined, setMicroservice: (microName: string) => void) {
    if (service === undefined) {
        return <div></div>
    }
    return (
        <div className="flex flex-wrap py-4 justify-center items-center">
            {service.microservices.map((micro: AreaMicroservices) =>
                <div key={`${micro.name}-${micro.ref_name}`} className="flex flex-row">
                    <MicroServiceCard
                        setMicroserviceAction={(): void => {
                            setMicroservice(micro.ref_name)
                        }}
                        microServicesColor={getColorForService(service.ref_name)}
                        title={micro.name}
                        description={`Service ${micro.ref_name}`}
                        microserviceName={micro.ref_name}
                    />
                </div>
            )}
        </div>
    )
}

export function renderIngredientsInput(
    ingredients: Map<string, Ingredient> | undefined,
    values: string[],
    setValues: React.Dispatch<React.SetStateAction<string[]>>
) {
    if (ingredients === undefined) {
        return <div></div>;
    }

    return (
        <>
            <div className="pt-3"></div>
            {Object.entries(ingredients).map(([ingredient, details] : [string, Ingredient], index: number) => (
                <div key={index} className="flex flex-col justify-center items-center">
                    <TooltipProvider>
                        <Tooltip>
                            <TooltipTrigger>
                                <p className="p-2 left-0 text-2xl text-white">
                                    {ingredient.charAt(0).toUpperCase() + ingredient.slice(1)}
                                </p>
                            </TooltipTrigger>
                            <TooltipContent className={"text-xl bg-white text-black font-bold border-4 border-black"}>
                                {details.description}
                            </TooltipContent>
                        </Tooltip>
                    </TooltipProvider>
                    <InputFieldComponent
                        ingredient={ingredient}
                        details={details}
                        index={index}
                        values={values}
                        setValues={setValues}
                    />
                </div>
            ))}
            <div className="pb-12"></div>
        </>
    );
}


const filterAreaByType = (services: AreaServices[], type: string) => {
    return services.filter((service: AreaServices): boolean => {
        return service.microservices.find((micro: AreaMicroservices): boolean => {
            return micro.type == type
        }) != undefined
    }).map((service: AreaServices) => {
        return {
            name: service.name,
            ref_name: service.ref_name,
            microservices: service.microservices.filter((micro: AreaMicroservices): boolean => {
                return micro.type == type
            })
        }
    })
}

// /!\ Disabling any error check because an ingredient can be of any type /!\
// eslint-disable-next-line
const convertIngredient = (ingredient: string | undefined, obj: Ingredient): any => {
    if (ingredient === undefined) {
        return null
    }

    switch (obj.type) {
        case "int":
            return parseInt(ingredient)
        case "float":
            return parseFloat(ingredient)
        case "bool":
            return ingredient.toLowerCase() === "true"
        case "time":
            return ""
        case "date":
            return ingredient
        default:
            return ingredient
    }
}

const filterServiceByRefName = (services: AreaServices[], refName: string): AreaServices | undefined => {
    return services.find((service: AreaServices): boolean => service.ref_name === refName)
}

export default function CreatePage({ services, uid }: { services: AreaServices[], uid: number }) {
    const [isTokenActionPresent, setTokenAction] = useState(true);
    const [isTokenReactionPresent, setTokenReaction] = useState(true);
    const router = useRouter();
    const { toast } = useToast()


    const actions: AreaServices[] = useMemo(() => {
        return filterAreaByType(services, "action")
    }, [services])

    const [actionName, setActionName] = React.useState("");
    const [microActionName, setMicroActionName] = React.useState("");

    const actionServiceChosen: AreaServices | undefined = useMemo((): AreaServices | undefined => {
        return filterServiceByRefName(actions, actionName)
    }, [actions, actionName])

    const reactions: AreaServices[] = useMemo(() => {
        return filterAreaByType(services, "reaction")
    }, [services])

    const [reactionName, setReactionName] = React.useState(""); // Later it will be an array of strings
    const [microReactionName, setMicroReactionName] = React.useState("");

    const reactionServiceChosen: AreaServices | undefined = useMemo((): AreaServices | undefined => {
        // Loop here with an array of reactionServiceChosen
        return filterServiceByRefName(reactions, reactionName)
    }, [reactions, reactionName])

    const [ingredientValuesActions, setIngredientValuesActions] = useState<string[]>([]);
    const [ingredientValuesReactions, setIngredientValuesReactions] = useState<string[]>([]);

    useEffect((): void => {
        if (actionName != "" && reactionName != "") {
            getUserTokens().then((res) => {
                let actionToken = false;
                if (actionName != "dt" && actionName != "weather") {
                    actionToken = res.includes(actionName);
                } else {
                    actionToken = true;
                }
                const reactionToken = res.includes(reactionName);
                setTokenAction(actionToken)
                setTokenReaction(reactionToken)
            }).catch((err) => {
                console.log(err)
            })
        }
    }, [actionName, reactionName])

    useEffect((): void => {
        setMicroActionName("")
    }, [actionName])

    useEffect((): void => {
        setMicroReactionName("")
    }, [reactionName])

    const handleSubmit = (formData: FormData): void => {
        if (actionServiceChosen === undefined || reactionServiceChosen === undefined) {
            console.error("Both action and reaction must be selected");
            return;
        }
        toast({
            title: "Area creation was sucessful",
            description: "Your new area is now running and available on this page.",
            variant: 'default',
            duration: 3000,
        })
        setTimeout((): void => {
            router.push("myareas/")
        }, 800);

        const payload: AreaCreateBody = {
            user_id: uid,
            action: {
                service: actionServiceChosen?.ref_name,
                microservice: microActionName,
                ingredients: {}
            },
            reactions: [{ // Here put all the reactions as an array
                service: reactionServiceChosen?.ref_name,
                microservice: microReactionName,
                ingredients: {}
            }]
        }

        const microAction: AreaMicroservices | undefined = actionServiceChosen.microservices.find((ms: AreaMicroservices): boolean => ms.ref_name === microActionName)
        const microReaction: AreaMicroservices | undefined = reactionServiceChosen.microservices.find((ms: AreaMicroservices): boolean => ms.ref_name === microReactionName)

        if (microAction === undefined || microReaction === undefined) {
            console.error("No microservice chosen.");
            return;
        }

        console.log("Micro action ing: ", microAction.ingredients)
        Object.entries(microAction.ingredients).forEach(([key, type]): void => {
            payload.action.ingredients[key] = convertIngredient(formData.get(key)?.toString(), type)
        })

        console.log("Micro reaction ing: ", microReaction.ingredients)
        Object.entries(microReaction.ingredients).forEach(([key, type]): void => {
            payload.reactions[0].ingredients[key] = convertIngredient(formData.get(key)?.toString(), type)
        })
        console.log(payload)
        create(payload).catch(error => { console.log(error) });
    };

    return (
        <Form action={handleSubmit}>
            <div className="pt-20 my-16 bg-white h-full w-full flex flex-col justify-center items-center p-8">
                <div
                    className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center"
                >
                    <MicroserviceCreateZone
                        services={actions} name={actionName} setNameAction={setActionName}
                        microServiceName={microActionName} serviceChosen={actionServiceChosen}
                        setServiceNameAction={setMicroActionName} ingredientsValues={ingredientValuesActions}
                        setIngredientValuesAction={setIngredientValuesActions} microServiceType={"action"}
                        textColor={"text-blue-500"}
                    />
                </div>
                <hr className="h-32 w-4 bg-gray-300" />
                <div // Later for multiple reactions this will be a loop
                    className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center"
                >
                    <MicroserviceCreateZone
                        services={reactions} name={reactionName} setNameAction={setReactionName}
                        microServiceName={microReactionName} serviceChosen={reactionServiceChosen}
                        setServiceNameAction={setMicroReactionName} ingredientsValues={ingredientValuesReactions}
                        setIngredientValuesAction={setIngredientValuesReactions} microServiceType={"reaction"}
                        textColor={"text-red-500"}
                    />
                </div>
                {(!isTokenActionPresent || !isTokenReactionPresent) && (
                    <div className={"font-bold mt-4 text-xl "}>
                        <p>You cannot create this area</p>
                        <p>There is no account linked to AREA for the following services </p>
                        {!isTokenActionPresent && (
                            <p className={"font-bold mx-4"}> Action : {actionName}</p>
                        )}
                        {!isTokenReactionPresent && (
                            <p className={"font-bold mx-4"}> Reaction :  {reactionName}</p>
                        )}
                    </div>
                )}


                <Button
                    type="submit"
                    className="mt-8 px-6 py-3 bg-green-500 text-white rounded-lg text-3xl font-bold"
                    disabled={microActionName === "" || microReactionName === "" || !isTokenActionPresent || !isTokenReactionPresent}
                >
                    Create AREA
                </Button>
            </div>
        </Form>
    );
}
