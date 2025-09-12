import Breadcrumbs from "@/components/Breadcrumbs";
import PollDetails from "@/components/PollDetails";
import { Poll } from "@/types/poll";
import { dateIsoToLocalString } from "@/utils/date";

async function fetchPoll(uuid: string): Promise<Poll | null> {
    const res = await fetch(
        `${process.env.API_URL ?? "http://localhost:8080"}/polls/${uuid}`,
        {
            next: { revalidate: 60 },
        }
    );
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
    params: Promise<{ uuid: string }>;
}) {
    const { uuid } = await params;
    const poll = await fetchPoll(uuid);
    if (!poll) return <div>Poll not found</div>;

    return (
        <>
            <Breadcrumbs />
            <PollDetails poll={poll} />
        </>
    );
}
