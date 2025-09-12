import Breadcrumbs from "@/components/Breadcrumbs";
import PollCard from "@/components/PollCard";
import { Poll } from "@/types/poll";
import { dateIsoToLocalString } from "@/utils/date";

async function fetchPolls(): Promise<Poll[]> {
    const res = await fetch(`http://localhost:8080/api/v1/polls`, {
        // cache control: cache for 60s; or use { cache: "no-store" }  to disable caching
        next: { revalidate: 60 },
    });

    if (!res.ok) {
        console.error("Failed to fetch polls", res.status);
        return [];
    }

    const data = (await res.json()) as Poll[];
    return data.map((r) => ({
        ...r,
        createdAt: dateIsoToLocalString(r.createdAt),
        updatedAt: dateIsoToLocalString(r.updatedAt),
        author: {
            ...r.author,
            createdAt: dateIsoToLocalString(r.author.createdAt),
            passwordChangedAt: dateIsoToLocalString(r.author.passwordChangedAt),
        },
    }));
}

export default async function Polls() {
    const polls = await fetchPolls();

    return (
        <main className="flex flex-1 flex-col">
            <Breadcrumbs />
            <section className="flex flex-1 flex-col items-center space-y-4 p-2">
                <hr className="border-border w-full" />
                <h2 className="text-lg font-medium">Available polls</h2>
                <hr className="border-border w-full" />

                <div className="grid grid-cols-2 gap-3">
                    {polls.map((poll) => (
                        <PollCard
                            key={poll.id}
                            poll={poll}
                        />
                    ))}
                </div>
            </section>
        </main>
    );
}
