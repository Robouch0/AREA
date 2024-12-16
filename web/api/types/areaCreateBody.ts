interface Service {
    service: string,
    microservice: string,
    ingredients: Map<string, any>,
}
interface AreaCreateBody {
    user_id: number,
    action: Service,
    reaction: Service,
}
