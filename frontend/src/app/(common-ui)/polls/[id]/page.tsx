import Breadcrumbs from "@/components/Breadcrumbs";
import PollDetails from "@/components/Poll/PollDetails";
import { Poll } from "@/types/poll";
import { dateIsoToLocalString } from "@/utils/date";

async function fetchPoll(id: string): Promise<Poll | null> {
    const res = await fetch(`http://localhost:8080/api/v1/polls/${id}`, {
        next: { revalidate: 60 },
    });
    if (!res.ok) return null;
    const raw = await res.json();
    return {
        ...raw,
        createdAt: dateIsoToLocalString(raw.createdAt),
        updatedAt: dateIsoToLocalString(raw.updatedAt),
    } as Poll;
}

export default async function PollPage({
    params,
}: {
    params: Promise<{ id: string }>;
}) {
    const { id } = await params;
    const poll = await fetchPoll(id);
    if (!poll) return <div>Poll not found</div>;

    return (
        <>
            <Breadcrumbs />
            <PollDetails poll={poll} />
        </>
    );
}
