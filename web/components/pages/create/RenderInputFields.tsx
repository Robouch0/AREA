import {Ingredient} from "@/api/types/areaStatus";
import * as React from "react";
import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/utils/thirdPartyComponents/shadcn/tooltip";
import {InputFieldComponent} from "@/components/pages/create/InputFieldComponent";

export function renderIngredientsInput(
    ingredients: Map<string, Ingredient> | undefined,
    values: string[],
    setValues: (values: string[]) => void,
    indexService: number
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
                                <div className={"flex flex-row items-center justify-center"}>
                                    <p className="p-2 left-0 text-2xl text-white">
                                        {ingredient.charAt(0).toUpperCase() + ingredient.slice(1)}
                                    </p>
                                    {details.required ? <p className={"text-lg text-white left-0"}> * </p> : null}
                                </div>
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
                        indexService={indexService}
                        values={values}
                        setValuesAction={setValues}
                    />
                </div>
            ))}
            <div className="pb-12"></div>
        </>
    );
}
