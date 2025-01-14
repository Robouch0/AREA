'use client';
import * as React from "react";
import {useEffect, useState, useMemo} from "react";
import {Button} from "@/components/ui/utils/Button";
import {create} from "@/api/createArea";
import Form from 'next/form';
import MicroserviceCreateZone from "@/components/ui/services/MicroserviceCreateZone";
import {convertIngredient} from "@/lib/utils";
import {AreaServices, AreaMicroservices} from "@/api/types/areaStatus";
import {AreaCreateBody} from "@/api/types/areaCreateBody";
import {getUserTokens} from "@/api/getUserInfos";
import {useRouter} from "next/navigation";
import {useToast} from "@/hooks/use-toast";
import {filterAreaByType, filterServiceByRefName} from "@/lib/filterCreateAreas";
import TokenStatus from "@/components/pages/create/TokenStatus";
import {AppRouterInstance} from "next/dist/shared/lib/app-router-context.shared-runtime";

export interface ServiceState {
    name: string;
    microServiceName: string;
    ingredientValues: string[];
}

interface TokenState {
    action: boolean;
    reactions: boolean[];
}

export default function CreatePage({services, uid}: { services: AreaServices[], uid: number }) {
    const [tokens, setTokens] = useState<TokenState>({action: true, reactions: []});
    const [action, setAction] = useState<ServiceState>({name: "", microServiceName: "", ingredientValues: []});
    const [reactions, setReactions] = useState<ServiceState[]>(
        [{name: "", microServiceName: "", ingredientValues: []}]);
    const router: AppRouterInstance = useRouter();
    const {toast} = useToast();

    const addReaction = () => {
        setReactions([...reactions, {name: "", microServiceName: "", ingredientValues: []}]);
    }

    const actions = useMemo(() => filterAreaByType(services, "action"), [services]);
    const reactionsList = useMemo(() => filterAreaByType(services, "reaction"), [services]);

    useEffect(() => {
        if (action.name || reactions.some(r => r.name)) {
            getUserTokens().then((res: string[]) => {
                const actionToken: boolean = action.name !== "dt" && action.name !== "weather" &&  action.name != "crypto" ? res.includes(
                    action.name) : true;
                const reactionTokens: boolean[] = reactions.map(r => res.includes(r.name))
                setTokens({action: actionToken, reactions: reactionTokens});
            }).catch((err) => {
                console.log(err)
            });
        }
    }, [action.name, reactions]);

    const handleSubmit = (formData: FormData): void => {
        if (!action.name || reactions.some(r => !r.name)) {
            console.error("Action and all reactions must be selected");
            return;
        }

        const payload: AreaCreateBody = {
            user_id: uid,
            action: {
                service: action.name,
                microservice: action.microServiceName,
                ingredients: {}
            },
            reactions: reactions.map(reaction => ({
                service: reaction.name,
                microservice: reaction.microServiceName,
                ingredients: {}
            }))
        };

        const actionService: AreaServices | undefined = filterServiceByRefName(actions, action.name);
        const actionMicroService: AreaMicroservices | undefined = actionService?.microservices.find(
            ms => ms.ref_name === action.microServiceName);
        if (actionMicroService) {
            Object.entries(actionMicroService.ingredients).forEach(([key, type]) => {
                payload.action.ingredients[key] = convertIngredient(formData.get(`-1-${key}`)?.toString(), type);
            });
        }

        reactions.forEach((reaction, index) => {
            const reactionService: AreaServices | undefined = filterServiceByRefName(reactionsList, reaction.name);
            const reactionMicroService: AreaMicroservices | undefined = reactionService?.microservices.find(
                ms => ms.ref_name === reaction.microServiceName);
            if (reactionMicroService) {
                Object.entries(reactionMicroService.ingredients).forEach(([key, type]) => {
                    payload.reactions[index].ingredients[key] = convertIngredient(
                        formData.get(`${index}-${key}`)?.toString(), type);
                });
            }
        })

        console.log(payload);
        create(payload).then(() => {
            toast({
                title: "Area creation was successful",
                description: "Your new area is now running and available on this page.",
                variant: 'default',
                duration: 3000,
            });
            setTimeout(() => router.push("myareas/"), 800);
        }).catch(console.error);
    };

    return (
        <Form action={handleSubmit}>
            <div className="pt-20 my-16 bg-white h-full w-full flex flex-col justify-center items-center p-8">
                <div
                    className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center"
                >
                    <MicroserviceCreateZone
                        index={-1}
                        serviceChosen={filterServiceByRefName(actions, action.name)}
                        services={actions}
                        name={action.name}
                        setNameAction={(name) => {
                            setAction(prev => ({...prev, name: name}))
                            setAction(prev => ({...prev, microServiceName: "", ingredientValues: []}))
                        }
                        }
                        microServiceName={action.microServiceName}
                        setServiceNameAction={(name) => {
                            setAction(prev => ({...prev, microServiceName: name}))
                            setAction(prev => ({...prev, ingredientValues: []}))
                        }
                        }
                        ingredientsValues={action.ingredientValues}
                        setIngredientValuesAction={(values) => setAction(prev => ({...prev, ingredientValues: values}))}
                        microServiceType={"action"}
                        textColor={"text-blue-500"}
                    />
                </div>
                {reactions.map((reaction, index) => (
                    <React.Fragment key={index}>
                    <hr className="h-32 w-4 bg-gray-300"/>
                        <div
                            className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center"
                        >
                            <MicroserviceCreateZone
                                index={index}
                                serviceChosen={filterServiceByRefName(reactionsList, reaction.name)}
                                services={reactionsList}
                                name={reaction.name}
                                setNameAction={(name) => {
                                    const newReactions = [...reactions];
                                    newReactions[index] = {
                                        ...newReactions[index],
                                        name,
                                        microServiceName: "",
                                        ingredientValues: []
                                    };
                                    setReactions(newReactions);
                                }}
                                microServiceName={reaction.microServiceName}
                                setServiceNameAction={(name) => {
                                    const newReactions = [...reactions];
                                    newReactions[index] = {
                                        ...newReactions[index],
                                        microServiceName: name,
                                        ingredientValues: []
                                    };
                                    setReactions(newReactions);
                                }}
                                ingredientsValues={reaction.ingredientValues}
                                setIngredientValuesAction={(values) => {
                                    const newReactions = [...reactions];
                                    newReactions[index] = {...newReactions[index], ingredientValues: values};
                                    setReactions(newReactions);
                                }}
                                microServiceType={"reaction"}
                                textColor={"text-red-500"}
                            />
                        </div>
                    </React.Fragment>
                ))
                }

                <Button
                    type="button"
                    onClick={addReaction}
                    className="mt-8 px-6 py-3 bg-blue-500 text-white rounded-lg text-3xl font-bold"
                >
                    Add Another Reaction
                </Button>

                <TokenStatus
                    isTokenActionPresent={tokens.action}
                    isTokenReactionPresent={tokens.reactions.every((token) => token)}
                    actionName={action.name}
                    reactionName={reactions.map(r => r.name).join(", ")}
                />

                <Button
                    type="submit"
                    className="mt-8 px-6 py-3 bg-green-500 text-white rounded-lg text-3xl font-bold"
                    disabled={action.microServiceName === "" || reactions.some(
                        r => r.microServiceName === "") || !tokens.action || !tokens.reactions.every((token) => token)}
                >
                    Create AREA
                </Button>
            </div>
        </Form>
    );
}
