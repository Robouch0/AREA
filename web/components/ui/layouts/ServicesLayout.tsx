"use client";
import Navbar, {User} from "@/components/ui/layouts/Navbar";
import {Button} from "@/components/ui/utils/Button";
import { useRouter } from 'next/navigation';

export default function ServicesLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {
    const user:User = {profilePicture: String("/areaLogo.png"), imgHeight: 60, imgWidth: 60};
    const router = useRouter();
    return (
        <div className="min-h-screen flex flex-col">
            <header className="fixed z-50 top-0 bg-white w-full h-24 flex flex-row justify-end shadow-sm">
                <div>
                    <Button
                        className="hidden sm:block sm:text-slate-800 hover:bg-transparent shadow-none bg-transparent absolute left-0 mx-6 text-4xl my-6 font-bold "
                        onClick={() => (router.push("/services/"))}
                    >
                        AREA
                    </Button>
                </div>
                <div>
                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-32 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={() => (router.push("/services/create/"))}
                    >
                        Create
                    </Button>
                </div>
                <div>
                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-32 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={() => (router.push("/services/"))}
                    >
                        Explore
                    </Button>
                </div>
                    <div className="">
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
                    <div>
                        <button
                            className="mx-4 text-amber-50 font-medium py-2 focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none focus-visible:p-2"
                            onClick={() => router.push('/services/contact')}
                        >
                            Nous contacter
                        </button>
                    </div>
                    <div>
                        <button
                            className="mx-4 text-amber-50 font-medium hover:cursor-pointer focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none focus-visible:p-2"
                            onClick={() => {
                                document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                                document.cookie = "UID=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                                router.push('/');
                            }}
                        >
                            Se déconnecter
                        </button>
                    </div>
                </div>
            </footer>
        </div>
    );
}
