'use client';
import { ComboboxDemo } from "@/components/ui/ComboboxDemo";
import { MicroServiceCard } from "@/components/ui/microserviceCard";
import * as React from "react";
import {useEffect, useState} from "react";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";





const getColorForService = (refName) => {
    const colors = {
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

function renderIngredientsInput(ingredients: any[], values: string[], setValues: React.Dispatch<React.SetStateAction<string[]>>) {
    if (values.length !== ingredients.length) {
        const newValues = Array(ingredients.length).fill('');
        setValues(newValues);
    }

    return (
        <>
            {ingredients.map((ingredient, index) => (
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


export default function CreatePage(response : any) {
    const {services, microservices } = processResponseData(Object.values(response));

    const [actionValue, setActionValue] = React.useState("");
    const [reactionValue, setReactionValue] = React.useState("");
    const [actionId, setActionId] = React.useState(-1);
    const [reactionId, setReactionId] = React.useState(-1);
    const [ingredientValuesActions, setIngredientValuesActions] = useState<string[]>([]);
    const [ingredientValuesReactions, setIngredientValuesReactions] = useState<string[]>([]);


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


    const handleSubmit = () => {
        if (actionId === -1 || reactionId === -1) {
            console.error("Both action and reaction must be selected");
            return;
        }

        const actionMicroservice = microservices
            .find(ms => ms.key === actionValue)
            ?.actions[actionId];
        const reactionMicroservice = microservices
            .find(ms => ms.key === reactionValue)
            ?.actions[reactionId];

        if (!actionMicroservice || !reactionMicroservice) {
            console.error("Microservices not found");
            return;
        }

        const actionIngredients = Object.keys(actionMicroservice.ingredients || {});
        const reactionIngredients = Object.keys(reactionMicroservice.ingredients || {});

        const payload = {
            action: {
                service: actionValue,
                microservice: actionMicroservice.title,
                ingredients: actionIngredients.reduce((acc, key, index) => {
                    acc[key] = ingredientValuesActions[index];
                    return acc;
                }, {})
            },
            reaction: {
                service: reactionValue,
                microservice: reactionMicroservice.title,
                ingredients: reactionIngredients.reduce((acc, key, index) => {
                    acc[key] = ingredientValuesReactions[index];
                    return acc;
                }, {})
            }
        };

        console.log(payload);
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
                        <form>
                            {renderIngredientsInput(
                                Object.keys(microservices.find((ms) => ms.key === actionValue)?.actions?.at(actionId)?.ingredients || {}),
                                ingredientValuesActions,
                                setIngredientValuesActions
                            )}
                        </form>
                    </>
                )}
            </div>
            <hr className="h-32 w-4 bg-gray-300"/>
            <div
                className="bg-slate-800 !opacity-100 text-6xl font-bold w-2/3 py-4 rounded-3xl flex flex-col justify-start items-center">
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
                        <form>
                            {renderIngredientsInput(
                                Object.keys(microservices.find((ms) => ms.key === reactionValue)?.actions?.at(reactionId)?.ingredients || {}),
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
                        disabled={actionId === -1 || reactionId === -1}
                    >
                        Create AREA
                    </Button>
        </div>
    );
}
