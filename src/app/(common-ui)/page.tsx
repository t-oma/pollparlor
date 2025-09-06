import Poll from "@/components/Poll";

export default function Home() {
    return (
        <main className="flex flex-1">
            <section className="relative flex flex-1 flex-col items-center px-4 py-16">
                <h2 className="absolute top-4 text-center text-xl font-medium">
                    <span className="sr-only">Poll</span>
                    Programming languages
                </h2>

                <Poll />
            </section>
        </main>
    );
}
