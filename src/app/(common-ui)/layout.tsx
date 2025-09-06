import Footer from "@/components/Footer";
import Header from "@/components/Header";

export default function CommonLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <>
            <Header />

            {children}

            <Footer />
        </>
    );
}
