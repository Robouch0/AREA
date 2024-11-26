export default function InAppLayout({children,}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
    <>
        {children}
        <footer>
            <h1> hello </h1>
        </footer>
    </>

    );
}
