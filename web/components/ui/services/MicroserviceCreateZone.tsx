"use client";
import {ComboboxDemo} from "@/components/ui/utils/ComboboxDemo";
import * as React from "react";
import {renderMicroservices} from "@/components/pages/create/RenderMicroServices";
import {AreaMicroservices, AreaServices} from "@/api/types/areaStatus";
import {renderIngredientsInput} from "@/components/pages/create/RenderInputFields";

export default function MicroserviceCreateZone({
        services,
        name,
        setNameAction,
        microServiceName,
        serviceChosen,
        setServiceNameAction,
        ingredientsValues,
        setIngredientValuesAction,
        microServiceType,
        textColor,
        index,
    }: {
        services: AreaServices[],
        name: string,
        setNameAction: (name: string) => void;
        microServiceName: string,
        serviceChosen: AreaServices | undefined,
        setServiceNameAction: (name: string) => void,
        ingredientsValues: string[],
        setIngredientValuesAction: (values: []) => void,
        microServiceType: string,
        textColor: string,
        index: number
    }
) {
    return (
        <>
            <div className={textColor}>
                <h1 className="my-2">{microServiceType.toUpperCase()}</h1>
            </div>
            <ComboboxDemo services={services} serviceName={name} setValueAction={setNameAction}/>
            {
                microServiceName === "" && name !== "" && (
                    <div className={textColor}>
                        <h1 className="p-6 text-5xl">  Select your {microServiceType}</h1>
                    </div>
                )
            }
            {
                microServiceName === "" ? (
                    renderMicroservices(
                        serviceChosen,
                        (microName: string): void => {
                            setServiceNameAction(microName)
                        }
                    )
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {serviceChosen && serviceChosen.microservices.find(
                                (ms: AreaMicroservices): boolean => ms.ref_name === microServiceName)?.name || microServiceName}
                        </div>
                        {serviceChosen && renderIngredientsInput(
                            serviceChosen.microservices.find(
                                (ms: AreaMicroservices): boolean => ms.ref_name === microServiceName)?.ingredients,
                            ingredientsValues,
                            setIngredientValuesAction,
                            index
                        )}
                    </>
                )
            }
        </>);
}
