"use client";
import { ChangeEvent, useState, useEffect, useCallback } from "react";
import { Ingredient } from "@/api/types/areaStatus";
import { TimePickerDemo } from "@/components/ui/utils/TimePicker";
import { Calendar } from "@/components/ui/calendar";
import { Input } from "@/components/ui/utils/Input";

export function InputFieldComponent({
    ingredient,
    details,
    index,
    values,
    setValuesAction
}: {
    ingredient: string;
    details: Ingredient;
    index: number;
    values: string[];
    setValuesAction: (values: []) => void;
}) {
    const [selectedDate, setSelectedDate] = useState<Date>(new Date());
    const [selectedTime, setSelectedTime] = useState<Date | undefined>(new Date());
    const [selectedDayAndMonth, setSelectedDayAndMonth] = useState<Date | undefined>(new Date());

    useEffect(() => {
        if (details.type === "date" && selectedDate) {
            setValuesAction(prevValues => {
                const newValues = [...prevValues];
                newValues[index] = selectedDate.toISOString();
                console.log(newValues[index])
                return newValues;
            });
        }
    }, [selectedDate, details.type, index, setValuesAction])

    const updateDate = useCallback((dateInput: Date) => {
        const date = new Date(selectedDate);
        date.setDate(dateInput.getDate());
        date.setMonth(dateInput.getMonth());
        date.setFullYear(dateInput.getFullYear());

        setSelectedDate(date)
    }, [selectedDate, setSelectedDate])

    const updateTime = useCallback((dateInput: Date) => {
        const time = new Date(selectedDate);
        time.setUTCHours(dateInput.getUTCHours() + 1);
        time.setUTCMinutes(dateInput.getUTCMinutes());

        setSelectedDate(time);
    }, [selectedDate, setSelectedDate])


    switch (details.type) {
        case "date":
            return (
                <div className="bg-white rounded-2xl flex flex-col items-center justify-center">
                    <Calendar
                        mode="single"
                        selected={selectedDayAndMonth}
                        onSelect={(date) => {
                            if (date == undefined) {
                                return
                            }
                            setSelectedDayAndMonth(date)
                            updateDate(date)
                        }}
                        className="bg-white text-xl font-bold rounded-md shadow"
                    />
                    <div className="mb-8">
                        <TimePickerDemo date={selectedTime} setDate={(newTime) => {
                            if (newTime == undefined) {
                                return
                            }
                            setSelectedTime(newTime)
                            updateTime(newTime)
                        }} />
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
                        setValuesAction(newValues);
                    }}
                    required
                />
            )
    }
}
