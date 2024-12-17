import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]): string {
  return twMerge(clsx(inputs))
}

export function getColorForService(refName: string):string {
    const colors: { [key: string]: string; } = {
        dt: "green",
        hf: "orange",
        github: "gray"
    };
    return colors[refName] || "blue";
}
