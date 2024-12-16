interface Service {
    service: string,
    microservice: string,
    ingredients: Ingredients<any>,
}
interface AreaCreateBody {
    user_id: number,
    action: Service,
    reaction: Service,
}
