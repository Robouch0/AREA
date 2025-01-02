"use client";
import Navbar, {User} from "@/components/ui/layouts/Navbar";
import {Button} from "@/components/ui/utils/Button";
import { useRouter } from 'next/navigation';
import {AppRouterInstance} from "next/dist/shared/lib/app-router-context.shared-runtime";
import React, { useState } from "react";
import { Menu } from 'lucide-react';

export default function ServicesLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {
    const user:User = {profilePicture: String("/areaLogo.png"), imgHeight: 60, imgWidth: 60};
    const router: AppRouterInstance = useRouter();
    const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

    return (
        <div className="min-h-screen flex flex-col">
            <header className="fixed z-50 top-0 bg-white w-full h-24 flex flex-row justify-end shadow-sm">
                <div>
                    <Button
                        className="hidden sm:block sm:text-slate-800 hover:bg-transparent shadow-none bg-transparent absolute left-0 mx-6 text-4xl my-6 font-bold "
                        onClick={(): void => (router.push("/services/"))}
                    >
                        AREA
                    </Button>
                </div>

                {/* Large display format */}
                <div className="hidden lg:flex">

                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={(): void => (router.push("/services/myareas/"))}
                    >
                        My Areas
                    </Button>
                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={(): void => (router.push("/services/create/"))}
                    >
                        Create
                    </Button>
                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={(): void => (router.push("/services/"))}
                    >
                        Explore
                    </Button>
                    <Navbar
                        prop={user}
                    />
                </div>

                {/* Small display format */}
                <div className="lg:hidden flex items-center mr-6">
                    <Button
                        className="p-4 my-5 mx-6 bg-black text-2xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                        onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
                    >
                        Menu
                        <Menu size={20} className="ml-2" />
                    </Button>
                    <div>
                        <Navbar prop={user} />
                    </div>
                </div>

                {/* Mobile Menu Dropdown */}
                {mobileMenuOpen && (
                    <div className="lg:hidden fixed top-24 right-6 bg-white shadow-md z-40 rounded-md overflow-hidden">
                        <div className="flex flex-col p-2">
                            <Button
                                className="p-4 my-1 bg-black text-xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                                onClick={(): void => {
                                    router.push("/services/myareas/");
                                    setMobileMenuOpen(false);
                                }}
                            >
                                My Areas
                            </Button>
                            <Button
                                className="p-4 my-1 bg-black text-xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                                onClick={(): void => {
                                    router.push("/services/create/");
                                    setMobileMenuOpen(false);
                                }}
                            >
                                Create
                            </Button>
                            <Button
                                className="p-4 my-1 bg-black text-xl text-white font-bold rounded-3xl h-12 w-36 hover:text-white hover:border-4 hover:border-black focus-visible:border-slate-500 focus-visible:border-8"
                                onClick={(): void => {
                                    router.push("/services/");
                                    setMobileMenuOpen(false);
                                }}
                            >
                                Explore
                            </Button>
                        </div>
                    </div>
                )}
            </header>

                <main className="grow">
                {children}
            </main>

            <footer className="bg-slate-800 w-full">
                <div className="container mx-auto px-4 py-8">
                    <h1 className="text-amber-50 font-semibold text-xl p-2 mb-4">AREA</h1>
                    <div className="flex flex-row space-x-8">
                        <div className="flex flex-col space-y-2">
                            <button
                                className="text-amber-50 font-medium px-4 focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none"
                                onClick={(): void => router.push('/services/profile')}
                            >
                                My Profile
                            </button>
                            <button
                                className="text-amber-50 font-medium px-4 focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none"
                                onClick={(): void => router.push('/services/contact')}
                            >
                                Contact Us
                            </button>
                        </div>
                        <div className="flex flex-col space-y-2">
                            <button
                                className="text-amber-50 font-medium px-4 focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none"
                                onClick={(): void => router.push('/services/myareas')}
                            >
                                My Areas
                            </button>
                            <button
                                className="text-amber-50 font-medium px-4 hover:cursor-pointer focus-visible:border-slate-500 focus-visible:border-4 rounded-full focus-visible:outline-none"
                                onClick={(): void => {
                                    document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                                    document.cookie = "UID=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
                                    router.push('/');
                                }}
                            >
                                Disconnect
                            </button>
                        </div>
                    </div>
                </div>
            </footer>

        </div>
    );
}
