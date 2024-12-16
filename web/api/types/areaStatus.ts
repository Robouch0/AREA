/** Type of ingredients possible. */
type IngredientPossible = "string" | "int" | "bool" | "time"

/**
 *  This is a representation of an Area Microservices.
 *
 *  @interface AreaMicroservices
 *
 *  @member {Ingredients} ingredients of the microservice
 */
interface AreaMicroservices {
    /** Name of the microservice */
    name: string,
    /** Name of the microservice server side */
    ref_name: string,

    /** Action or Reaction */
    type: string

    ingredients: Map<string, IngredientPossible>
}

/**
 *  This is a representation of an Area Service.
 *
 *  @interface AreaServices
 *
 *  @member {AreaMicroservices} Microservices associated with the area
 */
interface AreaServices {
    /** Name of the service */
    name: string,
    /** Name of the service server side */
    ref_name: string,

    microservices: AreaMicroservices[]
}
