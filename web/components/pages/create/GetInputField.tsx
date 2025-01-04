"use client";
import {ChangeEvent, useState} from "react";
import * as React from "react";
import {Ingredient} from "@/api/types/areaStatus";
import {TimePickerDemo} from "@/components/ui/utils/TimePicker";
import {Calendar} from "@/components/ui/calendar";
import {Input} from "@/components/ui/utils/Input";

export function GetInputField(
    ingredient: string,
    details: Ingredient,
    index: number,
    values: string[],
    setValues: React.Dispatch<React.SetStateAction<string[]>>
) {
    const [selectedTime, setSelectedTime] = useState<Date | undefined>(new Date());
    const [selectedDayAndMonth, setSelectedDayAndMonth] = useState<Date | undefined>(new Date());

    const updateValues = () => {
        if (selectedTime && selectedDayAndMonth) {
            const formattedDate = `${selectedDayAndMonth.getDate().toString().padStart(2,
                '0')}/${(selectedDayAndMonth.getMonth() + 1).toString().padStart(2,
                '0')}/${selectedTime.getHours().toString().padStart(2,
                '0')}:${selectedTime.getMinutes().toString().padStart(2, '0')}`;
            setValues(prevValues => {
                const newValues = [...prevValues];
                newValues[index] = formattedDate;
                return newValues;
            });
        }
    };

    switch(details.type) {
        case "date" :
        return (
            <div className="bg-white rounded-2xl flex flex-col items-center justify-center">
                <Calendar
                    mode="single"
                    selected={selectedDayAndMonth}
                    onSelect={(date) => {
                        setSelectedDayAndMonth(date)
                        updateValues()
                    }}
                    className="bg-white text-xl font-bold rounded-md shadow"
                />
                <div className="mb-8">
                    <TimePickerDemo date={selectedTime} setDate={(newDate) => {
                        setSelectedTime(newDate)
                        updateValues()
                    }}/>
                </div>
                <Input
                    type="hidden"
                    name={`${ingredient}`}
                    id={`text-${index}`}
                    className="!text-2xl !opacity-80 rounded-2xl bg-white font-extrabold border-4 focus:border-black w-2/3 p-4 h-14 placeholder:text-2xl placeholder:font-bold placeholder:opacity-60"
                    aria-label="text"
                    value={values[index] || ''}
                    required
                />
            </div>
        );
        default:
        return (
            <Input
                type={"text"}
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
        )
    }

}
