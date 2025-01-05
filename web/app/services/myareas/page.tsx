"use server";

import {listUserAreas} from "@/api/listAreas";
import {AreaServicesWithId} from "@/api/types/areaStatus";
import MyAreas from "@/components/pages/myareas/MyAreas";

export  default async function MyareasBackend() {
    try {
        const areas : AreaServicesWithId[] = await listUserAreas()
        return <MyAreas userAreas={areas}></MyAreas>
    } catch (err) {
        throw err;
    }
}
