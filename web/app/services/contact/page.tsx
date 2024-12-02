import Image from "next/image";

export default function Page() {
    return (
        <div className="flex flex-col items-center justify-center mt-16">
            <h1 className="mt-16 font-semibold text-3xl text-white"> Une proposition de nouvelles fonctionnalité ? </h1>
            <h2 className="font-semibold p-4 text-3xl text-white"> Un retour à faire sur notre site ? </h2>
            <p  className="font-medium text-white text-xl" > Contactez nous à : area-project@noj.fr</p>
            <Image
                className="my-12 opacity-60"
                src="/bg.png"
                alt="bg Social Networks"
                width="320"
                height="320"
            />
        </div>
    );
}
