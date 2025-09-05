export default function Home() {
    return (
        <div className="flex min-h-screen flex-col font-sans">
            <header className="flex items-center justify-between p-4 lg:p-8">
                <h1 className="text-3xl font-bold">Pollparlor</h1>
            </header>
            <main className="flex flex-1 p-4 py-24 lg:px-16 lg:py-16">
                <section className="flex flex-1 flex-col lg:flex-row">
                    <h2 className="sr-only">Poll</h2>
                    <div className="flex flex-1/2 border dark:border-zinc-500">
                        <button
                            type="button"
                            className="flex-1 cursor-pointer p-4 transition-colors dark:bg-zinc-900 dark:hover:bg-zinc-800"
                        ></button>
                    </div>
                    <div className="h-1 w-full lg:h-auto lg:w-1 dark:bg-zinc-500"></div>
                    <div className="flex flex-1/2 border dark:border-zinc-500">
                        <button
                            type="button"
                            className="flex-1 cursor-pointer p-4 dark:bg-zinc-900 dark:hover:bg-zinc-800"
                        ></button>
                    </div>
                </section>
            </main>
            <footer className="flex items-center justify-between p-4 lg:p-8">
                <p className="text-sm">&copy; 2025 Pollparlor</p>
            </footer>
        </div>
    );
}
