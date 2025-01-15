"use client";
import {ComboboxDemo} from "@/components/ui/utils/ComboboxDemo";
import * as React from "react";
import {renderMicroservices} from "@/components/pages/create/RenderMicroServices";
import {AreaMicroservices, AreaServices} from "@/api/types/areaStatus";
import {renderIngredientsInput} from "@/components/pages/create/RenderInputFields";
import {TokenState} from "@/app/services/create/page";
import {OauthButton} from "@/components/ui/services/OauthButton";

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
        tokens,
    }: {
        services: AreaServices[],
        name: string,
        setNameAction: (name: string) => void;
        microServiceName: string,
        serviceChosen: AreaServices | undefined,
        setServiceNameAction: (name: string) => void,
        ingredientsValues: string[],
        setIngredientValuesAction: (values: string[]) => void,
        microServiceType: string,
        textColor: string,
        index: number,
        tokens: TokenState[]
    }
) {
    const isOauthButtonVisible = tokens.find(
        (tokenState) => tokenState.providerName == serviceChosen?.ref_name)?.isTokenPresent || serviceChosen === undefined
    return (
        <>
            <>
                <ComboboxDemo services={services.sort((a,b) => a.name.localeCompare(b.name))} serviceName={name} setValueAction={setNameAction} />
                {
                    microServiceName === "" && name !== "" ? (
                        isOauthButtonVisible ? (
                            <div className={textColor}>
                                <h1 className="p-6 text-5xl"> Select your {microServiceType}</h1>
                            </div>
                        ) : (
                            <div className={textColor}>
                                <h1 className="p-6 text-5xl">Link your {serviceChosen?.name} account</h1>
                            </div>
                        )
                    ) : null
                }
            </>

            {
                isOauthButtonVisible ?
                    <></> :
                    <OauthButton
                        service={`${tokens.find(
                            (tokenState) => tokenState.providerName == serviceChosen?.ref_name)?.providerName}`}
                        location="create"
                        textButton={"Link"}
                        className="my-2 mr-6 lg:mr-2 font-bold text-2xl bg-green-600 w-24 focus-visible:border-8 focus-visible:border-black focus-visible:ring-0"
                        ServiceIcon={null}
                    />
            }

            {
                microServiceName === "" && isOauthButtonVisible ? (
                    renderMicroservices(
                        serviceChosen,
                        (microName: string): void => {
                            setServiceNameAction(microName)
                        }
                    )
                ) : null
            }

            {
                microServiceName !== "" && isOauthButtonVisible ? (
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
                ) : null
            }
        </>);
}
