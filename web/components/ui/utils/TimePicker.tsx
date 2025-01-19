"use client";

import * as React from "react";
import { Clock } from "lucide-react";
import { Label } from "@/components/ui/utils/thirdPartyComponents/shadcn/label";
import { TimePickerInput } from "./thirdPartyComponents/TimePickerInput";

interface TimePickerDemoProps {
    date: Date | undefined;
    setDate: (date: Date | undefined) => void;
}

export function TimePickerDemo({ date, setDate }: TimePickerDemoProps) {
    const minuteRef = React.useRef<HTMLInputElement>(null);
    const hourRef = React.useRef<HTMLInputElement>(null);

    return (
        <div className="flex items-center space-x-2 pt-3">
            <Clock className="h-5 w-5 text-gray-500 mt-4" />
            <div className="flex items-center space-x-2">
                <div className="flex flex-col items-center">
                    <Label htmlFor="hours" className="text-sm font-medium text-gray-700">
                        Hours
                    </Label>
                    <TimePickerInput
                        picker="hours"
                        date={date}
                        setDate={setDate}
                        ref={hourRef}
                        onRightFocus={() => minuteRef.current?.focus()}
                        className="w-14 text-center"
                    />
                </div>
                <span className="text-2xl text-gray-500 mt-4">:</span>
                <div className="flex flex-col items-center">
                    <Label htmlFor="minutes" className="text-sm font-medium text-gray-700">
                        Minutes
                    </Label>
                    <TimePickerInput
                        picker="minutes"
                        date={date}
                        setDate={setDate}
                        ref={minuteRef}
                        onLeftFocus={() => hourRef.current?.focus()}
                        className="w-14 text-center"
                    />
                </div>
            </div>
        </div>
    );
}
