import Breadcrumbs from "@/components/Breadcrumbs";
import PollsGrid from "@/components/Poll/PollsGrid";
import { fetchPolls } from "@/lib/api/polls";

export default async function Polls() {
    const polls = await fetchPolls();

    return (
        <main className="flex flex-1 flex-col">
            <Breadcrumbs />
            <section className="flex flex-1 flex-col items-center space-y-4 p-2">
                <hr className="border-border w-full" />
                <h2 className="text-lg font-medium">Available polls</h2>
                <hr className="border-border w-full" />

                <PollsGrid polls={polls} />
            </section>
        </main>
    );
}
