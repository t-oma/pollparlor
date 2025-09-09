import Poll from "@/components/Poll";

export default function Home() {
    return (
        <main className="flex flex-1">
            <section className="relative flex flex-1 flex-col items-center px-8 py-24">
                <div className="dark:border-border absolute top-2 rounded-md border p-2">
                    <h2 className="text-center text-xl font-medium">
                        <span className="sr-only">Poll</span>
                        Programming languages
                    </h2>
                </div>

                <Poll />
            </section>
        </main>
    );
}
