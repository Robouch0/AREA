import {AreaMicroservices, AreaServices} from "@/api/types/areaStatus";
import {MicroServiceCard} from "@/components/ui/services/areaCards/MicroserviceCard";
import {getColorForService} from "@/lib/utils";
import * as React from "react";

export  function renderMicroservices(service: AreaServices | undefined, setMicroservice: (microName: string) => void) {
    if (service === undefined) {
        return <div></div>
    }
    return (
        <div className="flex flex-wrap py-4 justify-center items-center">
            {service.microservices.map((micro: AreaMicroservices) =>
                <div key={`${micro.name}-${micro.ref_name}`} className="flex flex-row">
                    <MicroServiceCard
                        setMicroserviceAction={(): void => {
                            setMicroservice(micro.ref_name)
                        }}
                        microServicesColor={getColorForService(service.ref_name)}
                        title={micro.name}
                        description={`Service ${micro.ref_name}`}
                        microserviceName={micro.ref_name}
                    />
                </div>
            )}
        </div>
    )
}
