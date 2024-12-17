"use client";
import {ComboboxDemo} from "@/components/ui/utils/ComboboxDemo";
import * as React from "react";
import {renderIngredientsInput, renderMicroservices} from "@/components/pages/create/CreatePage";

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
        textColor
    }: {
        services: AreaServices[],
        name: string,
        setNameAction: React.Dispatch<React.SetStateAction<string>>,
        microServiceName: string,
        serviceChosen: AreaServices | undefined,
        setServiceNameAction: React.Dispatch<React.SetStateAction<string>>,
        ingredientsValues: string[],
        setIngredientValuesAction: React.Dispatch<React.SetStateAction<string[]>>,
        microServiceType: string,
        textColor: string
    }
) {
    return (
        <>
            <div className={textColor}>
                <h1 className="my-2">{microServiceType.toUpperCase()}</h1>
            </div>
            <ComboboxDemo services={services} serviceName={name} setValue={setNameAction}/>
            {

                microServiceName === "" && name !== "" && (
                    <div className={textColor}>
                        <h1 className="p-6 text-5xl">Veuillez sélectionner une {microServiceType}</h1>
                    </div>
                )
            }
            {
                microServiceName === "" ? (
                    renderMicroservices(
                        serviceChosen,
                        (microName) => {
                            setServiceNameAction(microName)
                        }
                    )
                ) : (
                    <>
                        <div className="p-2 my-4 text-xl flex flex-wrap text-white">
                            {serviceChosen && serviceChosen.microservices.find(
                                (ms) => ms.ref_name === microServiceName)?.name || microServiceName}
                        </div>
                        {serviceChosen && renderIngredientsInput(
                            serviceChosen.microservices.find(
                                (ms) => ms.ref_name === microServiceName)?.ingredients,
                            ingredientsValues,
                            setIngredientValuesAction
                        )}
                    </>
                )
            }
        </>);
}
