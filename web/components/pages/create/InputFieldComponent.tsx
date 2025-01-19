"use client";
import {ChangeEvent} from "react";
import {Ingredient} from "@/api/types/areaStatus";
import {Input} from "@/components/ui/utils/thirdPartyComponents/shadcn/Input";
import {CalendarTimeInput} from "@/components/pages/create/CalendarTimeInput";

export function InputFieldComponent({
    ingredient,
    details,
    index,
    values,
    setValuesAction,
    indexService
}: {
    ingredient: string;
    details: Ingredient;
    index: number;
    values: string[];
    setValuesAction: (values: string[]) => void;
    indexService: number;
}) {

    switch (details.type) {
        case "date":
            return (<CalendarTimeInput
                ingredient={ingredient} details={details} index={index} values={values} indexService={indexService}
                setValuesAction={(data: string) => {
                    const newValues: string[] = [...values]
                    newValues[index] = data;
                    setValuesAction(newValues);
                }}
            ></CalendarTimeInput>)
        case "int":
            return (
                <Input
                    type={"number"}
                    pattern="[0-9]+"
                    inputMode={"numeric"}
                    onKeyDown={(e) => {
                        const key = e.key;
                        const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab'];
                        if (!/[0-9]/.test(key) && !allowedKeys.includes(key)) {
                            e.preventDefault();
                        }
                    }}
                    name={`${indexService}-${ingredient}`}
                    id={`${indexService}-${ingredient}`}
                    className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60 focus-visible:!ring-0 shadow-none duration-200"
                    aria-label="number int input field"
                    value={values[index] || ''}
                    onChange={(e: ChangeEvent<HTMLInputElement>): void => {
                        const newValues: string[] = [...values];
                        newValues[index] = e.target.value;
                        setValuesAction(newValues);
                    }}
                    required={details.required || undefined}
                />
            )
        case "float":
            return (
                <Input
                    type={"number"}
                    inputMode={"decimal"}
                    pattern="[0-9]*\.?[0-9]*"
                    onKeyDown={(e) => {
                        const key = e.key;
                        const allowedKeys = ['Backspace', 'Delete', 'ArrowLeft', 'ArrowRight', 'Tab'];

                        if (!/[0-9.]/.test(key) && !allowedKeys.includes(key)) {
                            e.preventDefault();
                            return;
                        }
                    }}
                    name={`${indexService}-${ingredient}`}
                    id={`${indexService}-${ingredient}`}
                    className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60 focus-visible:!ring-0 shadow-none duration-200"
                    aria-label="float number input field"
                    value={values[index] || ''}
                    onChange={(e: ChangeEvent<HTMLInputElement>): void => {
                        const newValues: string[] = [...values];
                        newValues[index] = e.target.value;
                        setValuesAction(newValues);
                    }}
                    required={details.required || undefined}
                />

            )
        default:
            return (
                <>
                    <Input
                        type={"text"}
                        name={`${indexService}-${ingredient}`}
                        id={`${indexService}-${ingredient}`}
                        className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus-visible:border-black w-full p-4 h-16 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60 focus-visible:!ring-0 shadow-none duration-200"
                        aria-label="basic text input field"
                        value={values[index] || ''}
                        onChange={(e: ChangeEvent<HTMLInputElement>): void => {
                            const newValues: string[] = [...values];
                            newValues[index] = e.target.value;
                            setValuesAction(newValues);
                        }}
                        required={details.required || undefined}
                    />
                </>
            )
    }
}
