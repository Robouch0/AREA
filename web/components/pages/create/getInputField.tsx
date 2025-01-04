"use client";
import {ChangeEvent} from "react";
import {Input} from "@/components/ui/utils/Input";
import * as React from "react";
import {Ingredient} from "@/api/types/areaStatus";

export function getInputField(ingredient: string, details: Ingredient, index: number, values: string[],
    setValues: React.Dispatch<React.SetStateAction<string[]>>) {

    switch (details.type) {
        case "date":
            return (
                <>
                    <Input
                        type="time"
                        name={`${ingredient}`}
                        id={`text-${index}`}
                        className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                        aria-label="text"
                        value={values[index] || ''}
                        onChange={(e: ChangeEvent<HTMLInputElement>): void => {
                            const newValues: string[] = [...values];
                            newValues[index] = e.target.value;
                            setValues(newValues);
                        }}
                        required
                    />
                </>
            )
        default:
            return (
                <>
                    <Input
                        type="text"
                        name={`${ingredient}`}
                        id={`text-${index}`}
                        className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                        aria-label="text"
                        value={values[index] || ''}
                        onChange={(e: ChangeEvent<HTMLInputElement>): void => {
                            const newValues: string[] = [...values];
                            newValues[index] = e.target.value;
                            setValues(newValues);
                        }}
                        required
                    />
                </>
            )
    }

}
