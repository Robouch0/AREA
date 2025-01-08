import {AreaMicroservices, AreaServices} from "@/api/types/areaStatus";

export const filterAreaByType = (services: AreaServices[], type: string) => {
    return services.filter((service: AreaServices): boolean => {
        return service.microservices.find((micro: AreaMicroservices): boolean => {
            return micro.type == type
        }) != undefined
    }).map((service: AreaServices) => {
        return {
            name: service.name,
            ref_name: service.ref_name,
            microservices: service.microservices.filter((micro: AreaMicroservices): boolean => {
                return micro.type == type
            })
        }
    })
}

export const filterServiceByRefName = (services: AreaServices[], refName: string): AreaServices | undefined => {
    return services.find((service: AreaServices): boolean => service.ref_name === refName)
}
