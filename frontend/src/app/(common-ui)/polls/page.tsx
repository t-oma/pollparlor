import Breadcrumbs from "@/components/Breadcrumbs";
import Divider from "@/components/Divider";
import PollsGrid from "@/components/Poll/PollsGrid";
import { fetchPolls } from "@/lib/api/polls";

export default async function PollsPage() {
    const polls = await fetchPolls();

    return (
        <main className="flex flex-1 flex-col">
            <Breadcrumbs />
            <section className="flex flex-1 flex-col items-center space-y-4 p-2">
                <Divider />
                <h1 className="text-lg font-medium">Available polls</h1>
                <Divider />

                <PollsGrid polls={polls} />
            </section>
        </main>
    );
}
