import Navbar from "@/components/ui/navbar";
export default function InAppLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <div className="min-h-screen flex flex-col">
            <header className="bg-slate-800 w-full h-24">
                <Navbar prop={{profilePicture: String("/areaLogo.png")}}>

                </Navbar>
            </header>
            <main className="grow">
                {children}
            </main>
            <footer className="bg-slate-800 w-full">
                <div className="container mx-auto px-4 py-8">
                    <h1 className="text-amber-50 font-semibold text-xl mb-4">AREA</h1>
                    <a href="/services/contact" className="mx-4 text-amber-50 font-medium"> Nous contacter </a>
                </div>
            </footer>
        </div>
    );
}
