'use client';
import { ComboboxDemo } from "@/components/ui/ComboboxDemo";
import { MicroServiceCard } from "@/components/ui/microserviceCard";
import * as React from "react";
import {useEffect} from "react";

const getColorForService = (refName) => {
    const colors = {
        dt: "green",
        hf: "yellow",
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
        actions: item.microservices.map((ms, index) => ({
            id: index,
            title: ms.name,
            description: `Type: ${ms.type}, Ref: ${ms.ref_name}`,
            type: ms.type,
            ingredients: ms.ingredients
        }))
    }));

    return { services, microservices };
};

export default function CreatePage(response : any) {
    console.log(response)
    const {services, microservices } = processResponseData(Object.values(response));
    console.log(services);
    console.log(microservices);

    const [actionValue, setActionValue] = React.useState("");
    const [reactionValue, setReactionValue] = React.useState("");
    const [actionId, setActionId] = React.useState(-1);
    const [reactionId, setReactionId] = React.useState(-1);

    useEffect(() => {
        setActionId(-1);
    }, [actionValue]);

    useEffect(() => {
        setReactionId(-1);
    }, [reactionValue]);

    const renderMicroservices = (value, id, setId, type) => {
        const filteredMicroservices = microservices
            .find((ms) => ms.key === value)
            ?.actions.filter((action) => action.type === type);

        return (
            <div className="flex flex-wrap py-4 justify-center items-center">
                {filteredMicroservices?.map((action) => (
                    <div key={action.id} className="flex flex-row">
                        <MicroServiceCard
                            id={action.id}
                            setId={setId}
                            microServicesColor={microservices.find((ms) => ms.key === value)?.color}
                            title={action.title}
                            description={action.description}
                            service={microservices
                                .find((ms) => ms.key === value)
                                .key
                            }
                        />
                    </div>
                ))}
            </div>
        );
    };

    return (
        <div className="my-16 bg-white h-full w-full flex flex-col justify-center items-center p-8">
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                <h1 className="my-2 text-blue-500">ACTION</h1>
                <ComboboxDemo services={services} value={actionValue} setValue={setActionValue} />
                {actionValue !== "" && actionId === -1 && (
                    <h1 className="p-6 text-blue-500 text-5xl">Veuillez sélectionner une action</h1>
                )}
                {actionId === -1 ? (
                    renderMicroservices(actionValue, actionId, setActionId, "action")
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === actionValue)?.actions?.at(actionId)?.title}
                        </div>
                        {/*<div className="text-xl flex flex-wrap text-white">*/}
                        {/*    {microservices.find((ms) => ms.key === actionValue)?.actions?.at(actionId)?.description}*/}
                        {/*</div>*/}
                    </>
                )}
            </div>
            <hr className="h-32 w-4 bg-gray-300" />
            <div className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
                <h1 className="my-2 text-red-500">REACTION</h1>
                <ComboboxDemo services={services} value={reactionValue} setValue={setReactionValue} />
                {reactionValue !== "" && reactionId === -1 && (
                    <h1 className="p-6 text-red-500 text-5xl">Veuillez sélectionner une réaction</h1>
                )}
                {reactionId === -1 ? (
                    renderMicroservices(reactionValue, reactionId, setReactionId, "reaction")
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {microservices.find((ms) => ms.key === reactionValue)?.actions?.at(reactionId)?.title}
                        </div>
                        {/*<div className="text-xl flex flex-wrap text-white">*/}
                        {/*    {microservices.find((ms) => ms.key === reactionValue)?.actions?.at(reactionId)?.description}*/}
                        {/*</div>*/}
                    </>
                )}
            </div>
        </div>
    );
}
