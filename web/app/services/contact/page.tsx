import Image from "next/image";

export default function Contact() {
    return (
        <div className="pt-16 flex flex-col items-center justify-center mt-16">
            <h1 className="mt-16 font-semibold text-3xl text-black"> An idea of new feature ? </h1>
            <h2 className="font-semibold p-4 text-3xl text-black"> Some feedback to provide on our site ? </h2>
            <p  className="font-medium text-black text-xl" > Contact us : area-project@noj.fr</p>
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
