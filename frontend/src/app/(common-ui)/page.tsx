import { BookCopy } from "lucide-react";
import Link from "next/link";

export default function HomePage() {
    return (
        <main className="flex flex-1">
            <section className="relative flex flex-1 flex-col items-center px-8 py-24">
                <div className="dark:border-border absolute top-2 rounded-md border p-2">
                    <h2 className="text-center text-xl font-medium">
                        <span className="sr-only">Poll</span>
                        Programming languages
                    </h2>
                </div>

                <div className="flex flex-1 flex-col items-center justify-center">
                    <Link
                        href="/polls"
                        className="inline-flex items-center gap-4 rounded-md p-2 underline-offset-3 hover:underline"
                    >
                        <BookCopy className="h-5 w-5" />
                        <span className="text-xl">See polls page</span>
                    </Link>
                </div>
            </section>
        </main>
    );
}
