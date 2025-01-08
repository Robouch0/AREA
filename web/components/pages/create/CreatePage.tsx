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
    ingredientValues: [];
}

export default function CreatePage({services, uid}: { services: AreaServices[], uid: number }) {
    const [tokens, setTokens] = useState({ action: true, reaction: true });
    const [action, setAction] = useState<ServiceState>({ name: "", microServiceName: "", ingredientValues: [] });
    const [reaction, setReaction] = useState<ServiceState>({ name: "", microServiceName: "", ingredientValues: [] });
    const router: AppRouterInstance = useRouter();
    const { toast } = useToast();

    const actions = useMemo(() => filterAreaByType(services, "action"), [services]);
    const reactionsList = useMemo(() => filterAreaByType(services, "reaction"), [services]);

    useEffect(() => {
        if (action.name || reaction.name) {
            getUserTokens().then((res: string[]) => {
                const actionToken: boolean = action.name !== "dt" && action.name !== "weather" ? res.includes(action.name) : true;
                const reactionToken: boolean = res.includes(reaction.name);
                setTokens({ action: actionToken, reaction: reactionToken });
            }).catch((err) => {console.log(err)});
        }
    }, [action.name, reaction.name]);

        const actionService : AreaServices | undefined = filterServiceByRefName(actions, action.name);
        const actionMicroService: AreaMicroservices | undefined = actionService?.microservices.find(ms => ms.ref_name === action.microServiceName);

        const reactionService : AreaServices | undefined = filterServiceByRefName(reactionsList, reaction.name);
        const reactionMicroService : AreaMicroservices | undefined = reactionService?.microservices.find(ms => ms.ref_name === reaction.microServiceName);
    const handleSubmit = (formData: FormData): void => {
        if (!action.name || !reaction.name) {
            console.error("Both action and reaction must be selected");
            return;
        }
        if (actionMicroService) {
            Object.entries(actionMicroService.ingredients).forEach(([key, type]) => {
                payload.action.ingredients[key] = convertIngredient(formData.get(key)?.toString(), type);
            });
        }
        if (reactionMicroService) {
            Object.entries(reactionMicroService.ingredients).forEach(([key, type]) => {
                payload.reaction.ingredients[key] = convertIngredient(formData.get(key)?.toString(), type);
            });
        }

        const payload: AreaCreateBody = {
            user_id: uid,
            action: {
                service: action.name,
                microservice: action.microServiceName,
                ingredients: {}
            },
            reaction: {
                service: reaction.name,
                microservice: reaction.microServiceName,
                ingredients: {}
            }
        };

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
                <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                    <MicroserviceCreateZone
                        serviceChosen={actionService}
                        services={actions}
                        name={action.name}
                        setNameAction={(name) => setAction(prev => ({ ...prev, name: name }))}
                        microServiceName={action.microServiceName}
                        setServiceNameAction={(name) => setAction(prev => ({ ...prev, microServiceName: name }))}
                        ingredientsValues={action.ingredientValues}
                        setIngredientValuesAction={(values) => setAction(prev => ({ ...prev, ingredientValues: values }))}
                        microServiceType={"action"}
                        textColor={"text-blue-500"}
                    />
                </div>
                <hr className="h-32 w-4 bg-gray-300" />
                <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                    <MicroserviceCreateZone
                        serviceChosen={reactionService}
                        services={reactionsList}
                        name={reaction.name}
                        setNameAction={(name) => setReaction(prev => ({ ...prev, name }))}
                        microServiceName={reaction.microServiceName}
                        setServiceNameAction={(name) => setReaction(prev => ({ ...prev, microServiceName: name }))}
                        ingredientsValues={reaction.ingredientValues}
                        setIngredientValuesAction={(values) => setReaction(prev => ({ ...prev, ingredientValues: values }))}
                        microServiceType={"reaction"}
                        textColor={"text-red-500"}
                    />
                </div>

                <TokenStatus
                    isTokenActionPresent={tokens.action}
                    isTokenReactionPresent={tokens.reaction}
                    actionName={action.name}
                    reactionName={reaction.name}
                />

                <Button
                    type="submit"
                    className="mt-8 px-6 py-3 bg-green-500 text-white rounded-lg text-3xl font-bold"
                    disabled={action.microServiceName === "" || reaction.microServiceName === "" || !tokens.action || !tokens.reaction}
                >
                    Create AREA
                </Button>
            </div>
        </Form>
    );
}
