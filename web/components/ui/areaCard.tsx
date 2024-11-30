'use client';
import {Card, CardHeader, CardTitle, CardDescription} from "@/components/ui/card";
import Image from 'next/image'
import {Button} from "@/components/ui/button";
import { useRouter } from 'next/navigation';
export  function AreaCard({areaColor, category, action, reaction, areaTitle }) {
    const router = useRouter();

    const handleRedirection = () => {
        const textValue = "Hello world";

        router.push(`/services/create?textValue=${textValue}`);
        // ici on pourrait faire passer à la fontion le nom de l'action et de la reaction associé
        // comme ça quand le garts click on le met sur le /create et on remplis create avec les arguments dans l'url si y en a
    };
    return (
        <>
            <Button className="hover:bg-transparent shadow-none bg-transparent flex flex-row my-2 w-80 h-80"
                    onClick={handleRedirection}
            >
                <Card
                    className={"font-black text-3xl w-full h-full border-none hover:opacity-75 flex flex-col justify-between"}
                    style={{backgroundColor: areaColor}}
                >
                    <div className="flex flex-col flex-grow">
                        <div className="flex flex-row py-2 px-2">
                            <Image
                                className="m-2 object-contain"
                                src={action}
                                alt="img"
                                width="20"
                                height="20"
                            />
                            <Image
                                className="m-2 object-contain"
                                src={reaction}
                                alt="img"
                                width="20"
                                height="20"
                            />
                        </div>
                        <CardHeader className="text-wrap">
                            <CardTitle className="my-2 !text-2xl break-words text-white">{areaTitle}</CardTitle>
                        </CardHeader>
                    </div>
                        <div className="flex flex-row py-3 mx-4">
                            <Image
                                className=""
                                src="https://assets.ifttt.com/images/channels/28/icons/monochrome_regular.webp"
                                alt="img"
                                width="20"
                                height="20"
                            />
                            <CardDescription className="mx-2 text-black">
                                {category}
                            </CardDescription>
                        </div>
                </Card>

            </Button>
        </>
);
}
