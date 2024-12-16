interface Service {
    service: string,
    microservice: string,
    ingredients: Record<string, any>,
}
interface AreaCreateBody {
    user_id: number,
    action: Service,
    reaction: Service,
}
