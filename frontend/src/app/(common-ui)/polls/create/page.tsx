import Breadcrumbs from "@/components/Breadcrumbs";
import Divider from "@/components/Divider";
import CreatePollForm from "@/components/Poll/CreatePollForm";

export default function CreatePollPage() {
    return (
        <main className="flex flex-1 flex-col">
            <Breadcrumbs />

            <section className="flex flex-1 flex-col items-center space-y-4 p-2">
                <Divider />
                <h1 className="text-lg font-medium">Create poll</h1>
                <Divider />

                <CreatePollForm />
            </section>
        </main>
    );
}
