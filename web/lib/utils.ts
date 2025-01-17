import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import {Ingredient} from "@/api/types/areaStatus";

export function cn(...inputs: ClassValue[]): string {
  return twMerge(clsx(inputs))
}

export function getColorForService(refName: string):string {
    const colors: { [key: string]: string; } = {
        // dt: "#A76DAE", //  Purple
        // hf: "#D68A3A", //  Orange
        // github: "#808080", //  Gray
        // spotify: "#6BBF9A", //  Green
        // discord: "#A0A0A0", //  Gray
        // google: "#E2B800", //  Yellow (from Google)
    };

    return colors[refName] || "grey";
}

export function convertIngredient(ingredient: string | undefined, obj: Ingredient): string | boolean | null | number {
    if (ingredient === undefined) {
        return null
    }

    switch (obj.type) {
        case "int":
            return parseInt(ingredient)
        case "float":
            return parseFloat(ingredient)
        case "bool":
            return ingredient.toLowerCase() === "true"
        case "date":
            return ingredient
        default:
            return ingredient
    }
}
