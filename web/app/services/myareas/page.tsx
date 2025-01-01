"use server";

import {listUserAreas} from "@/api/listAreas";
import {AreaServicesWithId} from "@/api/types/areaStatus";
import Myareas from "@/components/pages/myareas/myareas";

export  default async function MyareasBackend() {
    try {
        const areas : AreaServicesWithId[] = await listUserAreas()
        console.log(areas)
        return <Myareas userAreas={areas}></Myareas>
    } catch (err) {
        throw err;
    }
}
