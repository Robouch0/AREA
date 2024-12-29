import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]): string {
  return twMerge(clsx(inputs))
}

export function getColorForService(refName: string):string {
    const colors: { [key: string]: string; } = {
        dt: "orange",
        hf: "orange",
        github: "gray",
        spotify: "green",
        discord: "gray",
        google: "red",
    };
    return colors[refName] || "blue";
}
