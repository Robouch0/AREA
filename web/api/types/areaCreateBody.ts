export interface Service {
    service: string,
    microservice: string,
    ingredients: Record<string, any>,
}
export interface AreaCreateBody {
    user_id: number,
    action: Service,
    reactions: Service[],
}
