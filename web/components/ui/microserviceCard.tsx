'use client';
import { Card, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ServiceIcon } from "@/components/ui/serviceIcon";

export function MicroServiceCard({ microServicesColor, title, description, setMicroservice, microserviceName }: { microServicesColor: string | undefined; title: string, description: string, setMicroservice: () => void, microserviceName: string }) {
    return (
        <>
            <Button
                className="hover:bg-transparent shadow-none bg-transparent my-2 w-80 h-64"
                onClick={() => setMicroservice()}
            >
                <Card
                    className={"font-black text-3xl w-full h-full border-none hover:opacity-75 items-center text-left py-2 flex flex-col justify-between"}
                    style={{ backgroundColor: microServicesColor }}
                >

                    <CardHeader className="text-wrap">
                        <ServiceIcon className="text-2xl text-white mb-8 mt-2" tag={microserviceName} />
                        <CardTitle className="my-2 !text-2xl break-words text-white">{title}</CardTitle>
                    </CardHeader>
                    <CardDescription className="px-6 py-4 text-black text-xl text-wrap">
                        {description}
                    </CardDescription>
                </Card>
            </Button>
        </>
    );
}
