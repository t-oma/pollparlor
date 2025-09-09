export default async function PollPage({
    params,
}: {
    params: Promise<{ uuid: string }>;
}) {
    const { uuid } = await params;

    return <div>Poll {uuid}</div>;
}
