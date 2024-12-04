'use client';
import Navbar, {User} from "@/components/ui/navbar";
import {Button} from "@/components/ui/button";
import {redirect} from "next/navigation";
import {checkAuthentification} from "@/api/backendCall";

export default function InAppLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {
    const user:User = {profilePicture: String("/areaLogo.png"), imgHeight: 60, imgWidth: 60};

    if (!checkAuthentification()) {
        redirect('/');
    }
    return (
        <div className="min-h-screen flex flex-col">
            <header className="bg-white w-full h-24 flex flex-row justify-end">
                <Button
                    className="hidden sm:block sm:text-slate-800 hover:bg-transparent shadow-none bg-transparent absolute left-0 mx-6 text-4xl my-6 font-bold"
                    onClick={() => (redirect("/services/"))}
                >
                    AREA
                </Button>
                <Button
                    className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-32 hover:text-white hover:border-4 hover:border-black"
                    onClick={() => (redirect("/services/create/"))}
                >
                    Create
                </Button>
                <Button
                    className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-32 hover:text-white hover:border-4 hover:border-black"
                    onClick={() => (redirect("/services/"))}
                >
                    Explore
                </Button>
                <div className="hover:">
                    <Navbar
                        prop={user}
                    >

                    </Navbar>
                </div>
            </header>

            <main className="grow">
                {children}
            </main>

            <footer className="bg-slate-800 w-full flex flex-row items-center justify-center">
                <div className="container mx-auto px-4 py-8 flex flex-col">
                    <h1 className="text-amber-50 font-semibold text-xl p-2 ">AREA</h1>
                    <a href="/services/contact" className="mx-4 text-amber-50 font-medium py-2 "> Nous contacter </a>
                    <a className="mx-4 text-amber-50 font-medium hover:cursor-pointer"
                       onClick={() =>  {
                        localStorage.removeItem("token");
                        redirect('/');
                        }}
                    > Se déconnecter </a>
                </div>
            </footer>
        </div>
    );
}
