'use client';
import {Card, CardHeader, CardTitle, CardDescription} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import { useRouter } from 'next/navigation';
import {ServiceIcon} from "@/components/ui/serviceIcon";
export  function AreaCard({areaColor, category, action, reaction, areaTitle}: {areaColor:string; category:string; action:string, reaction:string, areaTitle:string}) {
    const router = useRouter();

    const handleRedirection = () => {
        router.push(`/services/create?action=${action}&reaction=${reaction}&areatitle=${areaTitle}&color=${areaColor}`);
        // ici on pourrait faire passer à la fontion le nom de l'action et de la reaction associé
        // comme ça quand le garts click on le met sur le /create et on remplis create avec les arguments dans l'url si y en a
    };
    return (
        <>
            <Button className="hover:bg-transparent shadow-none bg-transparent flex flex-row my-2 w-80 h-80 focus-visible:border-slate-500 focus-visible:border-8 rounded-3xl focus-visible:outline-none"
                    onClick={handleRedirection}
            >
                <Card
                    className={"font-black text-3xl w-full h-full border-none hover:opacity-75 flex flex-col justify-between"}
                    style={{backgroundColor: areaColor}}
                >
                    <div className="flex flex-col">
                        <div className="flex flex-row py-2 px-2 my-2 items-center">

                            <ServiceIcon className="text-2xl text-white mx-2" tag={action}/>

                            <ServiceIcon className="text-2xl text-white" tag={reaction}/>
                        </div>
                        <CardHeader className="text-wrap">
                            <CardTitle className="my-2 !text-2xl break-words text-white">{areaTitle}</CardTitle>
                        </CardHeader>
                    </div>
                        <div className="flex flex-row py-3 mx-4">
                            <ServiceIcon tag={action}/>

                            <CardDescription className="mx-2 text-black">
                                {category}
                            </CardDescription>
                        </div>
                </Card>

            </Button>
        </>
);
}
