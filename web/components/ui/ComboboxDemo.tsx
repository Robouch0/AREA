"use client"

import * as React from "react"
import { Check, ChevronsUpDown } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList,
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from "@/components/ui/popover"


export function ComboboxDemo({ services, value, setValue }) {
    const [open, setOpen] = React.useState(false)

    return (
        <>
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild className="flex flex-row justify-between">
                <Button
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className="text-white hover:text-slate-200 bg-slate-800 border-2 hover:bg-slate-800 text-xl font-bold w-72"
                >
                    {value
                        ? services.find((service) => service.value === value)?.label
                        : "Sélectionnez un service"}
                    <ChevronsUpDown className="text-white h-4 shrink-0 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="text-black p-0 bg-slate-800 border-2 border-white">
                <Command className="bg-slate-800">
                    <CommandInput className="focus:!opacity-100 font-bold !text-white bg-slate-800 text-xl" placeholder="Cherchez un service..." />
                    <CommandList className="border-t-2 border-slate-700">
                    <CommandEmpty className="items-center justify-center mx-8 text-xl font-bold text-white">Aucun service trouvé.</CommandEmpty>
                        <CommandGroup>
                            {services.map((service) => (
                                <CommandItem
                                    key={service.value}
                                    value={service.value}
                                    className="font-bold text-xl text-white truncate"
                                    onSelect={(currentValue) => {
                                        setValue(currentValue === value ? "" : currentValue)
                                        setOpen(false)
                                    }}
                                >
                                    <Check
                                        className={cn(
                                            "mr-2 h-4 w-4",
                                            value === service.value ? "opacity-100" : "opacity-0"
                                        )}
                                    />
                                    {service.label}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    </>
    )
}
