"use client"

import * as React from "react"
import {Check, ChevronsUpDown} from "lucide-react"

import {cn} from "@/lib/utils"
import {Button} from "@/components/ui/utils/Button"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
} from "@/components/ui/utils/Command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/utils/Popover"
import {AreaServices} from "@/api/types/areaStatus";

export function ComboboxDemo({services, serviceName, setValueAction}: {
    services: AreaServices[];
    serviceName: string;
    setValueAction: React.Dispatch<React.SetStateAction<string>>
}) {
    const [open, setOpen] = React.useState(false)

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild className="flex flex-row justify-between">
                <Button
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className="text-white hover:text-slate-200 bg-slate-800 border-2 hover:bg-slate-800 text-xl font-bold w-72"
                >
                    {serviceName !== ""
                        ? services.find((service: AreaServices): boolean => service.ref_name === serviceName)?.ref_name
                        : "Select a Service"}
                    <ChevronsUpDown className="text-white h-4 shrink-0 opacity-50"/>
                </Button>
            </PopoverTrigger>
            <PopoverContent className="text-black p-0 bg-slate-800 border-2 border-white">
                <Command className="bg-slate-800">
                    <CommandInput
                        className="focus-visible:!opacity-100 font-bold !text-white bg-slate-800 text-xl"
                        placeholder="Search a service..."
                    />
                    <CommandList className="border-t-2 border-slate-700">
                        <CommandEmpty className="items-center justify-center mx-8 text-xl font-bold text-white">Aucun
                            service trouvé.</CommandEmpty>
                        <CommandGroup>
                            {services.map((service: AreaServices, i: number) => (
                                <CommandItem
                                    key={`${service.name}-${i}`}
                                    value={service.ref_name}
                                    className="font-bold text-xl text-white truncate"
                                    onSelect={(currentValue: string): void => {
                                        setValueAction(currentValue)
                                        setOpen(false)
                                    }}
                                >
                                    <Check
                                        className={cn(
                                            "mr-2 h-4 w-4",
                                            serviceName === service.ref_name ? "opacity-100" : "opacity-0"
                                        )}
                                    />
                                    {service.name}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}
