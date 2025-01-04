/** Type of ingredients possible. */
export type IngredientPossible = "string" | "int" | "bool" | "time" | "float"

export type Ingredient = {
    value: IngredientPossible,
    type: string,
    description: string,
    required: boolean,
}

/**
 *  This is a representation of an Area Microservices.
 *
 *  @interface AreaMicroservices
 *
 *  @member {Ingredients} ingredients of the microservice
 */
export interface AreaMicroservices {
    /** Name of the microservice */
    name: string,
    /** Name of the microservice server side */
    ref_name: string,

    /** Action or Reaction */
    type: string

    ingredients: Map<string, Ingredient>
}

/**
 *  This is a representation of an Area Service.
 *
 *  @interface AreaServices
 *
 *  @member {AreaMicroservices} Microservices associated with the area
 */
export interface AreaServices {
    /** Name of the service */
    name: string,
    /** Name of the service server side */
    ref_name: string,

    microservices: AreaMicroservices[]
}

export interface AreaServicesWithId extends AreaServices {
    /** Unique identifier for the user Area */
    id: string;
    /** Object containing the Area actions  */
    action: AreaServices;
    /** Object containing the Area reactions  */
    reactions: AreaServices[];
    /** Boolean to know if the area is active or not  */
    activated: boolean;

}
