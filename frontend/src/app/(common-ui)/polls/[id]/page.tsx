import Breadcrumbs from "@/components/Breadcrumbs";
import PollDetails from "@/components/Poll/PollDetails";
import { fetchPoll } from "@/lib/api/polls";

export default async function PollPage({
    params,
}: {
    params: Promise<{ id: string }>;
}) {
    const { id } = await params;
    const poll = await fetchPoll(id);
    if (!poll) return <div>Poll not found</div>;

    return (
        <main className="flex flex-1 flex-col">
            <Breadcrumbs />
            <PollDetails poll={poll} />
        </main>
    );
}
